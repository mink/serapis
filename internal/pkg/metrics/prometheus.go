package metrics

import (
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

	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe(":2112", nil)
}
