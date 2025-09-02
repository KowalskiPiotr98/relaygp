package server

import (
	smtp "github.com/emersion/go-smtp"
	log "github.com/sirupsen/logrus"
)

type backend struct{}

var _ smtp.Backend = (*backend)(nil)

func makeNewBackend() *backend {
	return &backend{}
}

func (b *backend) NewSession(c *smtp.Conn) (smtp.Session, error) {
	log.Debug("Setting up new session")
	return makeNewSession(), nil
}
