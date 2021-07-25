package incoming

import (
	"fmt"

	"github.com/gorilla/websocket"

	"serapis/internal/pkg/messages/outgoing"
	"serapis/internal/pkg/protocol"
)

type ReleaseVersionEvent struct {
	revision string
	client   string
}

func (e *ReleaseVersionEvent) Handle(conn *websocket.Conn) {
	fmt.Printf("ReleaseVersionEvent {revision: %s, client: %s}\n", e.revision, e.client)
}

type SecureLoginEvent struct {
	sso  string
}

func (e *SecureLoginEvent) Handle(conn *websocket.Conn) {
	fmt.Printf("SecureLoginEvent {sso: %s}\n", e.sso)

	// todo - authenticate

	err := conn.WriteMessage(websocket.BinaryMessage, protocol.NewOutgoingPacket(outgoing.SecureLoginOKComposer).All())
	if err != nil {
		go conn.Close()
	}
}

func init() {
	Events[4000] = func(packet *protocol.Packet) Event {
		return &ReleaseVersionEvent{
			revision: packet.ReadString(),
			client: packet.ReadString(),
		}
	}
	Events[2419] = func(packet *protocol.Packet) Event {
		return &SecureLoginEvent{
			sso:  packet.ReadString(),
		}
	}
}
