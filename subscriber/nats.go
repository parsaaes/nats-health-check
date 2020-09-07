package subscriber

import (
	"github.com/4lie/nats-health-check/config"
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

type Subscriber struct {
	conn    *nats.Conn
	handler Handler
}

func New(cfg config.NATS, h Handler) (*Subscriber, error) {
	opts := connectionOpts(cfg)

	c, err := nats.Connect(cfg.URL, opts...)
	if err != nil {
		return nil, err
	}

	return &Subscriber{
		conn:    c,
		handler: h,
	}, nil
}

func connectionOpts(cfg config.NATS) []nats.Option {
	var opts []nats.Option

	opts = append(opts, nats.ReconnectWait(cfg.ReconnectWait))

	opts = append(opts, nats.MaxReconnects(cfg.MaxReconnect))

	opts = append(opts, nats.ErrorHandler(func(_ *nats.Conn, _ *nats.Subscription, err error) {
		logrus.Errorf("Error: %s", err)
	}))

	opts = append(opts, nats.DisconnectErrHandler(func(_ *nats.Conn, err error) {
		logrus.Warnf("Disconnected: %s", err)
	}))

	opts = append(opts, nats.ReconnectHandler(func(nc *nats.Conn) {
		logrus.Infof("Reconnected [%s]", nc.ConnectedUrl())
	}))

	opts = append(opts, nats.ClosedHandler(func(nc *nats.Conn) {
		logrus.Warnf("Exiting: %v", nc.LastError())
	}))

	return opts
}

func (s *Subscriber) Subscribe() error {
	if _, err := s.conn.Subscribe(s.handler.Subject(), func(msg *nats.Msg) {
		s.handler.Handle(msg)
	}); err != nil {
		return err
	}

	return nil
}
