package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	addr              = flag.String("listen-address", ":9909", "The address to listen on for HTTP requests.")
)

var (
	testMetric = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "mike_test_metric",
		Help: "this is a test metric by mike",
		ConstLabels:prometheus.Labels{
			"Label1":"A test label",
		},
	})
)

func init() {
	prometheus.MustRegister(testMetric)
}

func main() {
	flag.Parse()

	go func() {
		for {
			testMetric.Add(1)
			time.Sleep(10 * time.Second)
		}
	}()

	fmt.Print("start")

	// Expose the registered metrics via HTTP.
	http.Handle("/metrics", promhttp.HandlerFor(
		prometheus.DefaultGatherer,
		promhttp.HandlerOpts{
		},
	))
	log.Fatal(http.ListenAndServe(*addr, nil))
}