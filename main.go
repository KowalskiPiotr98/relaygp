package main

import (
	"relaygp/client"
	"relaygp/config"
	"relaygp/server"

	log "github.com/sirupsen/logrus"
)

func init() {
	if config.GetCurrentConfig().DebugMode {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
}

func main() {
	if err := client.SendMessage(client.MessageOpts{}); err != nil {
		panic(err)
	}

	log.Infof("Setting up server...")
	server.Listen()
}
