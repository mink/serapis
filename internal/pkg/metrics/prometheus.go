package metrics

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var Gauges struct {
	Connections prometheus.Gauge
}

func Start() {
	Gauges.Connections = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "serapis",
		Name: "connections",
	})
	prometheus.MustRegister(Gauges.Connections)
	Gauges.Connections.Set(0)
	fmt.Println("Metrics -> Started Prometheus metrics")

	http.Handle("/metrics", promhttp.Handler())
	fmt.Println("Metrics -> Started Prometheus endpoint (0.0.0.0:2112)")
	go http.ListenAndServe(":2112", nil)
}
