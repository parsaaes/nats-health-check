package handler

import "github.com/nats-io/nats.go"

type Handler struct {
	subject string
	outputs []chan<- []byte
}

func New(subject string, outputs []chan<- []byte) *Handler {
	return &Handler{
		subject: subject,
		outputs: outputs,
	}
}

func (s *Handler) Handle(msg *nats.Msg) {
	for _, ch := range s.outputs {
		ch <- msg.Data
	}
}

func (s *Handler) Subject() string {
	return s.subject
}
