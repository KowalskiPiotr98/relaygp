package server

import (
	smtp "github.com/emersion/go-smtp"
	log "github.com/sirupsen/logrus"
)

type Handler func(from string, to string, data []byte) error

type backend struct {
	handlerFunc Handler
}

var _ smtp.Backend = (*backend)(nil)

func makeNewBackend(handlerFunc Handler) *backend {
	return &backend{
		handlerFunc: handlerFunc,
	}
}

func (b *backend) NewSession(c *smtp.Conn) (smtp.Session, error) {
	log.Debug("Setting up new session")
	return makeNewSession(b.handlerFunc), nil
}
