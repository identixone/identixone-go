package transport

import (
	"fmt"
)

type Transport uint8

func (t Transport) Validate() error {
	switch t {
	case Webhook, WebsocketClient, WebsocketServer:
		return nil
	default:
		return fmt.Errorf("unknown Transport %s", t)
	}
}

func (t Transport) String() string {
	switch t {
	case Webhook:
		return "webhook"
	case WebsocketClient:
		return "websocket client"
	case WebsocketServer:
		return "websocket server"
	default:
		return "unknown"
	}
}

const (
	Webhook Transport = iota
	WebsocketClient
	WebsocketServer
)
