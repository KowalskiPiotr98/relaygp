package server

import (
	"time"

	smtp "github.com/emersion/go-smtp"
	log "github.com/sirupsen/logrus"
)

func Listen() {
	s := smtp.NewServer(makeNewBackend())

	// test with: echo -ne 'ehlo localhost\r\nmail from:me@localhost\r\nrcpt to:you@localhost\r\ndata\r\ntest message\r\n.\r\n' | netcat localhost 9873
	s.Addr = "localhost:9873"
	s.Domain = "localhost"
	s.WriteTimeout = 10 * time.Second
	s.ReadTimeout = 10 * time.Second
	s.MaxMessageBytes = 1024 * 1024
	s.MaxRecipients = 50
	s.AllowInsecureAuth = true

	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
