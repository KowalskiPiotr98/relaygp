package main

import (
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
	log.Infof("Setting up server...")
	server.Listen()
}
