package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"serapis/internal/pkg/protocol"
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
		_, data, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Read error:", err)
			break
		}
		packet := protocol.NewPacket(data)
		fmt.Println("Packet received:", fmt.Sprintf("{header: %d, length: %d, bytes: %d}", packet.Header(), packet.Length(), packet.Data()))
	}
}
