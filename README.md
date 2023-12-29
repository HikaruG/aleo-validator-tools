# Simple SnarkOS monitoring tool

Grafana Monitoring tool for Aleo Validators.  
Only works by running it on Validator node.   

## Installation & Usage

### Prerequisites
Please make sure to have this: [aleo_simple_metrics](src/SimpleAleoMetrics.md)  
Install docker here: [docker](https://docs.docker.com/engine/install/ubuntu/)  


### Usage
If [node_exporter](https://github.com/prometheus/node_exporter) is not installed:  
`chmod +x installer_node_exporter.sh && ./installer_node_exporter.sh`

To start using this tool, simply run:  

```
git clone https://github.com/HikaruG/aleo-simple-monitoring.git 
cd aleo-simple-monitoring 
docker-compose up -d 
```

## What's included  

The below Aleo Logs Metrics SS was created using the Aleo Logs (Mock) Dashboard.  
More info about the Aleo Logs (Mock) [here](etc-promtail/README.md) 

Same for the Aleo Val Metrics, this was taken on a node running in --client mode.  

(Will Update the SS asap) 
- Aleo Logs Metrics<img width="1496" alt="Aleo_Logs_Mock" src="https://github.com/HikaruG/aleo-simple-monitoring/assets/43375172/39f2b234-42e7-4782-a637-f5b94453875e">

- Aleo Val Metrics<img width="1492" alt="Aleo_Val_Metrics" src="https://github.com/HikaruG/aleo-simple-monitoring/assets/43375172/a479351f-0501-47be-b93c-a21c27e3e8e1">

- Detailed Aleo node metrics [this one](https://grafana.com/grafana/dashboards/1860-node-exporter-full)
