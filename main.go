package main

import (
	"customize-exporter/files"
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	flag "github.com/spf13/pflag"
	"net/http"
	"time"
)

var appPort int

func init() {
	flag.IntVarP(&appPort, "app-port", "p", 9091, "customize exporter ")
}

func main() {
	flag.Parse()
	c := time.Tick(2 * time.Second)
	for range c {
		fmt.Println(files.Do())
	}

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(fmt.Sprintf(":%d", appPort), nil)
	fmt.Println("ip", appPort)
}
