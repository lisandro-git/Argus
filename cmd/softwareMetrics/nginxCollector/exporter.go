package nginxCollector

import (
	"argus/cmd"
	"argus/cmd/softwareMetrics/nginxCollector/client"
	collector2 "argus/cmd/softwareMetrics/nginxCollector/collector"
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"runtime/debug"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/model"
)

// positiveDuration is a wrapper of time.Duration to ensure only positive values are accepted
type positiveDuration struct{ time.Duration }

func (pd *positiveDuration) Set(s string) error {
	dur, err := parsePositiveDuration(s)
	if err != nil {
		return err
	}

	pd.Duration = dur.Duration
	return nil
}

func parsePositiveDuration(s string) (positiveDuration, error) {
	dur, err := time.ParseDuration(s)
	if err != nil {
		return positiveDuration{}, err
	}
	if dur < 0 {
		return positiveDuration{}, fmt.Errorf("negative duration %v is not valid", dur)
	}
	return positiveDuration{dur}, nil
}

func createPositiveDurationFlag(name string, value positiveDuration, usage string) *positiveDuration {
	flag.Var(&value, name, usage)
	return &value
}

type constLabel struct{ labels map[string]string }

func (cl *constLabel) Set(s string) error {
	labelList, err := parseConstLabels(s)
	if err != nil {
		return err
	}

	cl.labels = labelList.labels
	return nil
}

func (cl *constLabel) String() string {
	return fmt.Sprint(cl.labels)
}

func parseConstLabels(labels string) (constLabel, error) {
	if labels == "" {
		return constLabel{}, nil
	}

	constLabels := make(map[string]string)
	labelList := strings.Split(labels, ",")

	for _, l := range labelList {
		dat := strings.Split(l, "=")
		if len(dat) != 2 {
			return constLabel{}, fmt.Errorf("const label %s has wrong format. Example valid input 'labelName=labelValue'", l)
		}

		labelName := model.LabelName(dat[0])
		if !labelName.IsValid() {
			return constLabel{}, fmt.Errorf("const label %s has wrong format. %s contains invalid characters", l, labelName)
		}

		labelValue := model.LabelValue(dat[1])
		if !labelValue.IsValid() {
			return constLabel{}, fmt.Errorf("const label %s has wrong format. %s contains invalid characters", l, labelValue)
		}

		constLabels[dat[0]] = dat[1]
	}
	return constLabel{labels: constLabels}, nil
}

func createConstLabelsFlag(name string, value constLabel, usage string) *constLabel {
	flag.Var(&value, name, usage)
	return &value
}

func createClientWithRetries(getClient func() (interface{}, error), retries uint, retryInterval time.Duration) (interface{}, error) {
	var err error
	var nginxClient interface{}

	for i := 0; i <= int(retries); i++ {
		nginxClient, err = getClient()
		if err == nil {
			return nginxClient, nil
		}
		if i < int(retries) {
			log.Printf("Could not create Nginx Client. Retrying in %v...", retryInterval)
			time.Sleep(retryInterval)
		}
	}
	return nil, err
}

func parseUnixSocketAddress(address string) (string, string, error) {
	addressParts := strings.Split(address, ":")
	addressPartsLength := len(addressParts)

	if addressPartsLength > 3 || addressPartsLength < 1 {
		return "", "", fmt.Errorf("address for unix domain socket has wrong format")
	}

	unixSocketPath := addressParts[1]
	requestPath := ""
	if addressPartsLength == 3 {
		requestPath = addressParts[2]
	}
	return unixSocketPath, requestPath, nil
}

var (
	// Set during go build
	version string

	// Defaults values
	defaultListenAddress      = getEnv("LISTEN_ADDRESS", ":8080")
	defaultSecuredMetrics     = getEnvBool("SECURED_METRICS", false)
	defaultSslServerCert      = getEnv("SSL_SERVER_CERT", "")
	defaultSslServerKey       = getEnv("SSL_SERVER_KEY", "")
	defaultMetricsPath        = getEnv("TELEMETRY_PATH", "/metrics")
	defaultScrapeURI          = getEnv("SCRAPE_URI", "http://127.0.0.1/stub_status")
	defaultSslVerify          = getEnvBool("SSL_VERIFY", true)
	defaultSslCaCert          = getEnv("SSL_CA_CERT", "")
	defaultSslClientCert      = getEnv("SSL_CLIENT_CERT", "")
	defaultSslClientKey       = getEnv("SSL_CLIENT_KEY", "")
	defaultTimeout            = getEnvPositiveDuration("TIMEOUT", time.Second*5)
	defaultNginxRetries       = getEnvUint("NGINX_RETRIES", 0)
	defaultNginxRetryInterval = getEnvPositiveDuration("NGINX_RETRY_INTERVAL", time.Second*5)
	defaultConstLabels        = getEnvConstLabels("CONST_LABELS", map[string]string{})

	// Command-line flags
	listenAddr = flag.String("web.listen-address",
		defaultListenAddress,
		"An address or unix domain socket path to listen on for web interface and telemetry. The default value can be overwritten by LISTEN_ADDRESS environment variable.")
	securedMetrics = flag.Bool("web.secured-metrics",
		defaultSecuredMetrics,
		"Expose metrics using https. The default value can be overwritten by SECURED_METRICS variable.")
	sslServerCert = flag.String("web.ssl-server-cert",
		defaultSslServerCert,
		"Path to the PEM encoded certificate for the nginx-exporter metrics server(when web.secured-metrics=true). The default value can be overwritten by SSL_SERVER_CERT variable.")
	sslServerKey = flag.String("web.ssl-server-key",
		defaultSslServerKey,
		"Path to the PEM encoded key for the nginx-exporter metrics server (when web.secured-metrics=true). The default value can be overwritten by SSL_SERVER_KEY variable.")
	metricsPath = flag.String("web.telemetry-path",
		defaultMetricsPath,
		"A path under which to expose metrics. The default value can be overwritten by TELEMETRY_PATH environment variable.")
	scrapeURI = flag.String("nginx.scrape-uri",
		defaultScrapeURI,
		`A URI or unix domain socket path for scraping NGINX or NGINX Plus metrics.
For NGINX, the stub_status page must be available through the URI. For NGINX Plus -- the API. The default value can be overwritten by SCRAPE_URI environment variable.`)
	sslVerify = flag.Bool("nginx.ssl-verify",
		defaultSslVerify,
		"Perform SSL certificate verification. The default value can be overwritten by SSL_VERIFY environment variable.")
	sslCaCert = flag.String("nginx.ssl-ca-cert",
		defaultSslCaCert,
		"Path to the PEM encoded CA certificate file used to validate the servers SSL certificate. The default value can be overwritten by SSL_CA_CERT environment variable.")
	sslClientCert = flag.String("nginx.ssl-client-cert",
		defaultSslClientCert,
		"Path to the PEM encoded client certificate file to use when connecting to the server. The default value can be overwritten by SSL_CLIENT_CERT environment variable.")
	sslClientKey = flag.String("nginx.ssl-client-key",
		defaultSslClientKey,
		"Path to the PEM encoded client certificate key file to use when connecting to the server. The default value can be overwritten by SSL_CLIENT_KEY environment variable.")
	nginxRetries = flag.Uint("nginx.retries",
		defaultNginxRetries,
		"A number of retries the exporter will make on start to connect to the NGINX stub_status page/NGINX Plus API before exiting with an error. The default value can be overwritten by NGINX_RETRIES environment variable.")
	displayVersion = flag.Bool("version",
		false,
		"Display the NGINX exporter version.")

	// Custom command-line flags
	timeout = createPositiveDurationFlag("nginx.timeout",
		defaultTimeout,
		"A timeout for scraping metrics from NGINX or NGINX Plus. The default value can be overwritten by TIMEOUT environment variable.")

	nginxRetryInterval = createPositiveDurationFlag("nginx.retry-interval",
		defaultNginxRetryInterval,
		"An interval between retries to connect to the NGINX stub_status page/NGINX Plus API on start. The default value can be overwritten by NGINX_RETRY_INTERVAL environment variable.")

	constLabels = createConstLabelsFlag("prometheus.const-labels",
		defaultConstLabels,
		"A comma separated list of constant labels that will be used in every metric. Format is label1=value1,label2=value2... The default value can be overwritten by CONST_LABELS environment variable.")
)

func Start() {
	commitHash, commitTime, dirtyBuild := getBuildInfo()
	arch := fmt.Sprintf("%v/%v", runtime.GOOS, runtime.GOARCH)

	fmt.Printf("NGINX Prometheus Exporter version=%v commit=%v date=%v, dirty=%v, arch=%v, go=%v\n", version, commitHash, commitTime, dirtyBuild, arch, runtime.Version())

	if *displayVersion {
		os.Exit(0)
	}

	log.Printf("Starting...")

	registry := cmd.Registry

	buildInfoMetric := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "nginxexporter_build_info",
			Help: "Exporter build information",
			ConstLabels: collector2.MergeLabels(
				constLabels.labels,
				prometheus.Labels{
					"version": version,
					"commit":  commitHash,
					"date":    commitTime,
					"dirty":   strconv.FormatBool(dirtyBuild),
					"arch":    arch,
					"go":      runtime.Version(),
				},
			),
		},
	)
	buildInfoMetric.Set(1)

	registry.MustRegister(buildInfoMetric)

	// #nosec G402
	sslConfig := &tls.Config{InsecureSkipVerify: !*sslVerify}
	if *sslCaCert != "" {
		caCert, err := os.ReadFile(*sslCaCert)
		if err != nil {
			log.Fatalf("Loading CA cert failed: %v", err)
		}
		sslCaCertPool := x509.NewCertPool()
		ok := sslCaCertPool.AppendCertsFromPEM(caCert)
		if !ok {
			log.Fatal("Parsing CA cert file failed.")
		}
		sslConfig.RootCAs = sslCaCertPool
	}

	if *sslClientCert != "" && *sslClientKey != "" {
		clientCert, err := tls.LoadX509KeyPair(*sslClientCert, *sslClientKey)
		if err != nil {
			log.Fatalf("Loading client certificate failed: %v", err)
		}
		sslConfig.Certificates = []tls.Certificate{clientCert}
	}

	transport := &http.Transport{
		TLSClientConfig: sslConfig,
	}
	if strings.HasPrefix(*scrapeURI, "unix:") {
		socketPath, requestPath, err := parseUnixSocketAddress(*scrapeURI)
		if err != nil {
			log.Fatalf("Parsing unix domain socket scrape address %s failed: %v", *scrapeURI, err)
		}

		transport.DialContext = func(_ context.Context, _, _ string) (net.Conn, error) {
			return net.Dial("unix", socketPath)
		}
		newScrapeURI := "http://unix" + requestPath
		scrapeURI = &newScrapeURI
	}

	userAgent := fmt.Sprintf("NGINX-Prometheus-Exporter/v%v", version)
	userAgentRT := &userAgentRoundTripper{
		agent: userAgent,
		rt:    transport,
	}

	httpClient := &http.Client{
		Timeout:   timeout.Duration,
		Transport: userAgentRT,
	}

	srv := http.Server{
		ReadHeaderTimeout: 5 * time.Second,
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		log.Printf("Signal received: %v. Exiting...", <-signalChan)
		err := srv.Close()
		if err != nil {
			log.Fatalf("Error occurred while closing the server: %v", err)
		}
		os.Exit(0)
	}()

	ossClient, err := createClientWithRetries(func() (interface{}, error) {
		return client.NewNginxClient(httpClient, *scrapeURI)
	}, *nginxRetries, nginxRetryInterval.Duration)
	if err != nil {
		log.Fatalf("Could not create Nginx Client: %v", err)
	}
	cmd.RegisterCollector(collector2.NewNginxCollector(ossClient.(*client.NginxClient), "nginxplus", constLabels.labels))

	//if *securedMetrics {
	//	_, err = os.Stat(*sslServerCert)
	//	if err != nil {
	//		log.Fatalf("Cert file is not set, not readable or non-existent. Make sure you set -web.ssl-server-cert when starting your exporter with -web.secured-metrics=true: %v", err)
	//	}
	//	_, err = os.Stat(*sslServerKey)
	//	if err != nil {
	//		log.Fatalf("Key file is not set, not readable or non-existent. Make sure you set -web.ssl-server-key when starting your exporter with -web.secured-metrics=true: %v", err)
	//	}
	//	log.Printf("NGINX Prometheus Exporter has successfully started using https")
	//	log.Fatal(srv.ServeTLS(listener, *sslServerCert, *sslServerKey))
	//}

	log.Printf("NGINX Prometheus Exporter has successfully started")
}

type userAgentRoundTripper struct {
	agent string
	rt    http.RoundTripper
}

func (rt *userAgentRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	req = cloneRequest(req)
	req.Header.Set("User-Agent", rt.agent)
	return rt.rt.RoundTrip(req)
}

func cloneRequest(req *http.Request) *http.Request {
	r := new(http.Request)
	*r = *req // shallow clone

	// deep copy headers
	r.Header = make(http.Header, len(req.Header))
	for key, values := range req.Header {
		newValues := make([]string, len(values))
		copy(newValues, values)
		r.Header[key] = newValues
	}
	return r
}

func getBuildInfo() (string, string, bool) {
	var commitHash, commitTime string
	var dirtyBuild bool
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "", "", false
	}
	for _, kv := range info.Settings {
		switch kv.Key {
		case "vcs.revision":
			commitHash = kv.Value
		case "vcs.time":
			commitTime = kv.Value
		case "vcs.modified":
			dirtyBuild = kv.Value == "true"
		}
	}
	return commitHash, commitTime, dirtyBuild
}
