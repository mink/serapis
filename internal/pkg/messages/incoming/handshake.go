package incoming

import (
	"fmt"

	"serapis/internal/pkg/protocol"
)

type ReleaseVersionEvent struct {
	revision string
	client   string
}

func (e *ReleaseVersionEvent) Handle() {
	fmt.Println("Revision:", e.revision, "Client:", e.client)
}

func init() {
	Events[4000] = func(packet *protocol.Packet) Event {
		return &ReleaseVersionEvent{
			revision: packet.ReadString(),
			client: packet.ReadString(),
		}
	}
}
