package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))
)

func main() {
	opsProcessedCounter := promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_ops_counter",
		Help: "Counter",
		ConstLabels: prometheus.Labels{
			"host": fmt.Sprintf("%s-%d", os.Getenv("host"), rand.Int()),
		},
	})
	opsProcessedGauge := promauto.NewGauge(prometheus.GaugeOpts{
		Name: "myapp_processed_ops_gauge",
		Help: "Gauge",
		ConstLabels: prometheus.Labels{
			"host": fmt.Sprintf("%s-%d", os.Getenv("host"), rand.Int()),
		},
	})
	go func() {
		i := 1.0
		for {
			opsProcessedCounter.Inc()
			opsProcessedGauge.Set(i)
			i = i + 1
			time.Sleep(2 * time.Second)
		}
	}()
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("port")), nil)
	fmt.Println("Exiting?")
}
