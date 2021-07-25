package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func main() {
	fmt.Println("Serapis")

	go metrics()
	serve()
}

func metrics() {
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}

func serve() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":2096", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	for {
		_, packet, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf("read err - %s\n", err)
			break
		}
		fmt.Println(len(packet), "bytes received:", packet)
	}
}
