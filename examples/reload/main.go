package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var addr = flag.String("listen-address", ":8080", "The address to listen on for HTTP requests.")

var (
	simpleCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "simple_counter_metric",
		Help: "Simple counter metric",
	})
)

func main() {
	flag.Parse()
	prometheus.MustRegister(simpleCounter)

	// Periodically increment counter
	go func() {
		for {
			simpleCounter.Inc()
			time.Sleep(100 * time.Millisecond)
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(*addr, nil))
}
