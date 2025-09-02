package server

import (
	"io"

	smtp "github.com/emersion/go-smtp"
	log "github.com/sirupsen/logrus"
)

type session struct{}

var _ smtp.Session = (*session)(nil)

func makeNewSession() *session {
	return &session{}
}

func (s *session) Data(r io.Reader) error {
	log.Debug("Reading data")
	data, err := io.ReadAll(r)
	if err != nil {
		log.Warn(err)
		return err
	}
	log.Debugf("Data read: %s", (string)(data))
	return nil
}

func (s *session) Logout() error {
	log.Debug("Logging out")
	return nil
}

func (s *session) Mail(from string, opts *smtp.MailOptions) error {
	log.Debugf("Reading mail from: %s", from)
	return nil
}

func (s *session) Rcpt(to string, opts *smtp.RcptOptions) error {
	log.Debugf("Reading rcpt to: %s", to)
	return nil
}

func (s *session) Reset() {
	log.Debug("Resetting session")
}
