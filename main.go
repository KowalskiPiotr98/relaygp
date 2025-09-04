package main

import (
	"relaygp/client"
	"relaygp/config"
	"relaygp/server"

	log "github.com/sirupsen/logrus"
)

var (
	handler server.Handler = func(from, to string, data []byte) error {
		return client.SendMessage(client.MessageOpts{
			To:   to,
			Data: data,
		})
	}
)

func init() {
	if config.GetCurrentConfig().DebugMode {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
}

func main() {
	log.Infof("Setting up server...")
	server.Listen(handler)
}
