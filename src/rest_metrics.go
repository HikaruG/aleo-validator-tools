package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	localHeight = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "node_local_height",
		Help: "Local height of Aleo node",
	})

	networkHeight = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "node_network_height",
		Help: "Network height of Aleo node",
	})

	validatorStake = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "validator_total_stake",
		Help: "Total staked amount of the validator",
	})

	//for more context: https://github.com/AleoHQ/snarkVM/blob/b7c5f49ba0a6b573f5a1f6850338507152827f8c/ledger/committee/src/lib.rs#L44
	validatorIsOpen = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "validator_is_open",
		Help: "The bonding state of the validator",
	})
)

func init() {
	prometheus.MustRegister(localHeight)
	prometheus.MustRegister(networkHeight)
	prometheus.MustRegister(validatorStake)
	prometheus.MustRegister(validatorIsOpen)
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "RPC Query to Prometheus Metric exporter.\n\n")
		flag.PrintDefaults()
	}

	validatorAddress := flag.String("validator", "aleo1yzlta2q5h8t0fqe0v6dyh9mtv4aggd53fgzr068jvplqhvqsnvzq7pj2ke", "The validator address")
	listenPort := flag.String("listen.port", "9380", "port to listen on")
	localEndpoint := flag.String("local_endpoint", "http://localhost:3033", "local endpoint to connect to")
	publicEndpoint := flag.String("public_endpoint", "https://api.explorer.aleo.org/v1", "public endpoint to connect to")
	queryInterval := flag.Int("query.interval", 15, "interval in (s) between each query")
	flag.Parse()

	http.Handle("/metrics", promhttp.Handler())

	go func() {
		client := &http.Client{}
		for {
			updateLocalMetrics(client, *localEndpoint)
			updatePublicMetrics(client, *publicEndpoint, *validatorAddress)
			time.Sleep(time.Duration(*queryInterval) * time.Second)
		}
	}()

	fmt.Printf("Aleo Node Exporter started on port %s\n", *listenPort)
	http.ListenAndServe(":"+*listenPort, nil)
}

func updateLocalMetrics(client *http.Client, localEndpoint string) {
	local, err := getHeight(client, localEndpoint)
	if err != nil {
		fmt.Println("Error getting metrics:", err)
		return
	}
	localHeight.Set(float64(local))
}

func updatePublicMetrics(client *http.Client, publicEndpoint string, validatorAddress string) {
	network, err := getHeight(client, publicEndpoint)
	if err != nil {
		fmt.Println("Error getting metrics:", err)
		return
	}
	stake_value, bonding_status, err := getStakeStatus(publicEndpoint, validatorAddress)
	networkHeight.Set(float64(network))
	validatorStake.Set(float64(stake_value))
	if bonding_status {
		validatorIsOpen.Set(1)
	} else {
		validatorIsOpen.Set(0)
	}
}

func getHeight(client *http.Client, endpoint string) (int, error) {
	request := endpoint + "/testnet3/block/height/latest"
	resp, err := http.Get(request)
	if err != nil {
		fmt.Printf("error making request: %v", err)
		return 0, err
	}
	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return 0, err
	}
	height, err := strconv.Atoi(string(respBytes))
	if err != nil {
		fmt.Printf("Error converting height to int: %v\n", err)
		return 0, err
	}
	return height, err
}

func getStakeStatus(endpoint string, validatorAddress string) (uint64, bool, error) {
	request := endpoint + "/testnet3/committee/latest"
	resp, err := http.Get(request)
	if err != nil {
		return 0, false, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, false, fmt.Errorf("error reading response body: %v", err)
	}
	var response LatestCommittee
	err = json.Unmarshal(body, &response)
	if err != nil {
		return 0, false, err
	}
	values, ok := response.Members[validatorAddress]
	if !ok {
		return 0, false, fmt.Errorf("validatorAddress not found")
	}
	stake_value, ok1 := values[0].(float64)
	bonding_status, ok2 := values[1].(bool)
	if !ok1 || !ok2 {
		return 0, false, fmt.Errorf("invalid data types for address values")
	}
	return uint64(stake_value), bonding_status, nil
}
