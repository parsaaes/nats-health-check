package subscriber

import "github.com/nats-io/nats.go"

type Handler interface {
	Handle(msg *nats.Msg)
	Subject() string
}
