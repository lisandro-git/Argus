# Current Grafana Dashboard
![grafana_dashboard](img/grafana_dashboard.png)

---
# Table of Contents
- [Current Grafana Dashboard](#current-grafana-dashboard)
- [Table of Contents](#table-of-contents)
- [Metrics Collected](#metrics-collected)
- - [Hardware Metrics](#hardware-metrics)
- - [Network Metrics](#network-metrics)
- - [Software Metrics](#software-metrics)
- - - [Nginx](#Nginx)
- - - [Gitea](#Gitea)
- - [Operating System Metrics](#os-metrics)

---
# Metrics collected
| Metric                   | Description                           |
|--------------------------|---------------------------------------|
| Hardware Metrics         | Data related to all hardware devices  |
| Network Metrics          | Data related to all network devices   |
| Software Metrics         | Data related to all software packages |
| Operating System Metrics | Data related to the operating system  |

---
## Hardware Metrics
### CPU
| **Metric**    | **Description** | **Type** | **Unit** |
|---------------|-----------------|----------|----------|
| **CPUUser**   | CPU user        | Gauge    | %        |
| **CPUNice**   | CPU Nice        | Gauge    | %        |
| **CPUSys**    | CPU idle        | Gauge    | %        |
| **CPUIntr**   | CPU Interrupts  | Gauge    | %        |
| **CPUIdle**   | CPU CPU idle    | Gauge    | %        |
| **CPUStates** | CPU States      | Gauge    | %        |
| **CPUTemp**   | CPU States      | Counter  | %        |

### Memory
| **Metric** | **Description**       | **Type** | **Unit** |
|------------|-----------------------|----------|----------|
| **Total**  | Total memory          | Gauge    | Bytes    |
| **Free**   | Available free memory | Gauge    | Bytes    |
| **Usage**  | Memory Used           | Gauge    | Bytes    |

### Disk
| **Metric**      | **Description** | **Type** | **Unit** |
|-----------------|-----------------|----------|----------|
| **Total**       | Total disk      | Gauge    | Bytes    |
| **Free**        | Free disk       | Gauge    | Bytes    |
| **Usage**       | Disk Used       | Gauge    | Bytes    |
| **Percentage ** | Disk Percentage | Gauge    | Bytes    |

---
## Network Metrics
| **Metric**           | **Description**            | **Type** | **Unit** |
|----------------------|----------------------------|----------|----------|
| **networkLatency**   | Network latency            | Gauge    | ms       |
| **connectedClients** | Connected client on subnet | Gauge    | Count    |

---
## Software Metrics
### Nginx
| **Metric** | **Description**         | **Type** | **Unit** |
|------------|-------------------------|----------|----------|
| **Active** | Active Connections      | Gauge    | Count    |
| **Accepts**| Connection Accepted     | Gauge    | Count    |
| **Handled**| Connection Handled      | Gauge    | Count    |
| **Reading**| Request Headers read    | Gauge    | Count    |
| **Writing**| Response back to client | Gauge    | Count    |
| **Waiting**| Idle client connection  | Gauge    | Count    |
| **Total**  | Total Http requests     | Gauge    | Count    |

### Gitea
| **Metric**           | **Description**                          | **Type** | **Unit** |
|----------------------|------------------------------------------|----------|----------|
| **Active**           | Active                                   | Gauge    | Count    |
| **gitea_repository** | Gitea Repository (email, name, owner...) | %        | %        |

---
## OS Metrics
### Uptime
| **Metric** | **Description** | **Type** | **Unit** |
|------------|-----------------|----------|----------|
| **Uptime** | System Uptime   | Gauge    | Count    |
---
