package main

import (
	"relaygp/server"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.DebugLevel)
	log.Infof("Setting up server...")
	server.Listen()
}
