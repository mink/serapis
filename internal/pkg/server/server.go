package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"

	"serapis/internal/pkg/messages/incoming"
	"serapis/internal/pkg/metrics"
	"serapis/internal/pkg/network"
	"serapis/internal/pkg/protocol"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func Start() {
	http.HandleFunc("/", handler)
	fmt.Println("Server  -> Started TCP server (0.0.0.0:2096)")
	http.ListenAndServe(":2096", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		panic(err)
	}

	conn := network.NewConnection(ws)

	metrics.Gauges.Connections.Inc()

	defer func() {
		metrics.Gauges.Connections.Dec()
		go conn.Close()
	}()

	for {
		data, err := conn.Read()
		if err != nil {
			fmt.Println("Read error:", err)
			break
		}
		packet := protocol.NewInboundPacket(data)
		fmt.Println("Packet received:", fmt.Sprintf("{header: %d, length: %d, bytes: %d}", packet.Header(), packet.Length(), packet.Data()))

		event := incoming.Events[int(packet.Header())]
		if event != nil {
			event(packet).Handle(conn)
		} else {
			fmt.Println("Packet", packet.Header(), "invalid")
		}
	}
}
