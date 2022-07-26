package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	flag "github.com/spf13/pflag"
	"net/http"
)

var appPort int

func init() {
	flag.IntVar(&appPort, "app port", 9091, "customize exporter ")
}

func main() {
	flag.Parse()
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(fmt.Sprintf(":%d", appPort), nil)
	fmt.Println("ip", appPort)
}
