package main

import (
	"CRUD/src"
	"CRUD/src/infrastructure"
	"github.com/prometheus/client_golang/prometheus"
)

func init() {
	// we need to register the counter so prometheus can collect this metric
	prometheus.MustRegister(userStatus)
}
func main() {
	var server = src.Server{}
	infrastructure.Init()
	server.RunServer("3000")
	//c := exec.Command("top")
	//c.Stdin = os.Stdin
	//c.Stdout = os.Stdout
	//c.Stderr = os.Stderr
	//c.Run()
}

var userStatus = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_request_get_user_status_count", // metric name
		Help: "Count of status returned by user.",
	},
	[]string{"user", "status"}, // labels
)
