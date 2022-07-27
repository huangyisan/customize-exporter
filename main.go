package main

import (
	"customize-exporter/files"
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	flag "github.com/spf13/pflag"
	"net/http"
)

var (
	appPort int
	Version string
	Build   string
	version bool
	// exporter mode
	mode string
)

func init() {
	flag.BoolVarP(&version, "version", "v", false, "show app version")
	flag.IntVarP(&appPort, "app-port", "p", 9091, "customize exporter ")
}

func main() {
	flag.Parse()

	command()

	switch mode {
	case "file":
		files.Do()
	default:
		panic("模式错误")
	}

	http.Handle("/metrics", promhttp.Handler())
	fmt.Printf("Listen on :%d", appPort)
	http.ListenAndServe(fmt.Sprintf(":%d", appPort), nil)
}
