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
- - - [Nginx Metrics](#nginx)
- - - [Gitea Metrics](#gitea)
- - - [GitHub Metrics](#github)
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
| **Percentage** | Disk Percentage | Gauge    | Bytes    |

---
## Network Metrics
| **Metric**           | **Description**            | **Type** | **Unit** |
|----------------------|----------------------------|----------|----------|
| **networkLatency**   | Network latency            | Gauge    | ms       |
| **connectedClients** | Connected client on subnet | Gauge    | Count    |

---
## Software Metrics
### Nginx
| **Metric**  | **Description**         | **Type** | **Unit** |
|-------------|-------------------------|----------|----------|
| **Active**  | Active Connections      | Gauge    | Count    |
| **Accepts** | Connection Accepted     | Gauge    | Count    |
| **Handled** | Connection Handled      | Gauge    | Count    |
| **Reading** | Request Headers read    | Gauge    | Count    |
| **Writing** | Response back to client | Gauge    | Count    |
| **Waiting** | Idle client connection  | Gauge    | Count    |
| **Total**   | Total Http requests     | Gauge    | Count    |

### Gitea
| **Metric**   | **Description**                 | **Type** | **Unit** |
|--------------|---------------------------------|----------|----------|
| **Active**   | State of the Gitea service      | Gauge    | Count    |
| **ID**       | ID of the repository            | label    | %        |
| **Name**     | Name of the repo                | label    | %        |
| **Owner**    | Owner of the current repository | label    | %        |
| **Email**    | Email of the owner of the repo  | label    | %        |
| **RepoSize** | Size of the repository          | label    | %        |

### Github
| **Metric**        | **Description**            | **Type** | **Unit** |
|-------------------|----------------------------|----------|----------|
| **ID**            | ID of the repo             | label    | %        |
| **Name**          | Name of the repo           | label    | %        |
| **Owner**         | Owner of the current repo  | label    | %        |
| **RepoSize**      | Size of the repo           | label    | %        |
| **DefaultBranch** | Default branch of the repo | label    | %        |
| **CloneUrl**      | Url to clone the repo      | label    | %        |
| **Language**      | Main language of the repo  | label    | %        |
| **Description**   | Description of the repo    | label    | %        |
| **Visibility**    | Visibility of the repo     | label    | %        |
| **CreationDate**  | Creation date of the repo  | label    | %        |
| **LastUpdate**    | Last update of the repo    | label    | %        |

---
## OS Metrics
### Uptime
| **Metric** | **Description** | **Type** | **Unit** |
|------------|-----------------|----------|----------|
| **Uptime** | System Uptime   | Gauge    | Count    |
---
