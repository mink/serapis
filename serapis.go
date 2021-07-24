package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	fmt.Println("Serapis")

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
