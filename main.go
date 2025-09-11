package main

import (
	"bufio"
	"bytes"
	"relaygp/client"
	"relaygp/config"
	"relaygp/pgp"
	"relaygp/server"
	"slices"

	log "github.com/sirupsen/logrus"
)

var (
	handler server.Handler = func(from, to string, data []byte) error {
		reader := bufio.NewScanner(bytes.NewReader(data))
		headers := make([][]byte, 0)
		body := make([][]byte, 0)
		readingHeaders := true
		for reader.Scan() {
			line := reader.Text()
			if !readingHeaders {
				body = append(body, []byte(line))
				continue
			}
			headers = append(headers, []byte(line))
			if line == "" {
				readingHeaders = false
			}
		}
		message, err := pgp.Encrypt(bytes.Join(body, []byte("\n")))
		if err != nil {
			return err
		}

		return client.SendMessage(client.MessageOpts{
			To:   to,
			Data: slices.Concat(bytes.Join(headers, []byte("\n")), []byte("\n"), message),
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
