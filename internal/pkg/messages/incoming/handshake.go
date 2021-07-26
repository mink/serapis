package incoming

import (
	"fmt"
	"serapis/internal/pkg/game"
	"serapis/internal/pkg/network"

	"serapis/internal/pkg/messages/outgoing"
	"serapis/internal/pkg/protocol"
)

type ReleaseVersionEvent struct {
	revision string
	client   string
}

func (e *ReleaseVersionEvent) Handle(conn *network.Connection) {
	fmt.Printf("ReleaseVersionEvent {revision: %s, client: %s}\n", e.revision, e.client)
}

type SecureLoginEvent struct {
	sso  string
}

func (e *SecureLoginEvent) Handle(conn *network.Connection) {
	fmt.Printf("SecureLoginEvent {sso: %s}\n", e.sso)

	var user *game.User

	for i := range game.Users {
		if game.Users[i].SSO == e.sso {
			user = game.Users[i]
		}
	}

	if user == nil {
		go conn.Close()
		return
	}

	conn.User = user
	err := conn.Write(&outgoing.SecureLoginOKComposer{})
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
