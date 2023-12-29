# Simple SnarkOS simple monitoring tool

Currently, SnarkOS does not have any prometheus metrics exported by default.  
This monitoring tool uses the SnarkOS REST API and push prometheus metrics.  

## What's included
4 metrics exposed: 
- node_local_height
- node_network_height
- validator_total_stake
- validator_is_open

## Usecase
This tool works best when both the local endpoint and public endpoint are given.  
Example: alert based on the difference between the two heights, or when the total_stake does not increase.    

## Installation & Usage

### Binary
simple_aleo_metrics binary can be found in the src folder  
Compiled using go1.21.5 linux/amd64

### build from source
Requires [go](https://go.dev/doc/install) 
(tested on 1.20 and 1.21)
```
git clone https://github.com/HikaruG/aleo-simple-monitoring.git 
cd src
go build -o simple_aleo_metrics
./simple_aleo_metrics -<put_your_aleo_val_address>
```

### Parameters
Please set the validator flag, the default address won't work.  
```
./simple_aleo_metrics --help
RPC Query to Prometheus Metric exporter.

  -listen.port string
    	port to listen on (default "9380")
  -local_endpoint string
    	local endpoint to connect to (default "http://localhost:3033")
  -public_endpoint string
    	public endpoint to connect to (default "https://api.explorer.aleo.org/v1")
  -query.interval int
    	interval in (s) between each query (default 15)
  -validator string
    	The validator address (default "aleo1yzlta2q5h8t0fqe0v6dyh9mtv4aggd53fgzr068jvplqhvqsnvzq7pj2ke")
```
