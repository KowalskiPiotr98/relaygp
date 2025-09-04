package server

import (
	"io"

	smtp "github.com/emersion/go-smtp"
	log "github.com/sirupsen/logrus"
)

type session struct {
	data []byte
	from string
	to   string

	handlerFunc Handler
}

var _ smtp.Session = (*session)(nil)

func makeNewSession(handlerFunc Handler) *session {
	return &session{
		handlerFunc: handlerFunc,
	}
}

func (s *session) Data(r io.Reader) error {
	log.Debug("Reading data")
	data, err := io.ReadAll(r)
	if err != nil {
		log.Warn(err)
		return err
	}
	log.Debugf("Data read: %s", (string)(data))
	s.data = data

	if s.from == "" || s.to == "" {
		//todo: this is likely an error of some kind
		return nil
	}

	return s.handlerFunc(s.from, s.to, s.data)
}

func (s *session) Logout() error {
	log.Debug("Logging out")
	return nil
}

func (s *session) Mail(from string, opts *smtp.MailOptions) error {
	log.Debugf("Reading mail from: %s", from)
	s.from = from
	return nil
}

func (s *session) Rcpt(to string, opts *smtp.RcptOptions) error {
	log.Debugf("Reading rcpt to: %s", to)
	s.to = to
	return nil
}

func (s *session) Reset() {
	log.Debug("Resetting session")
}
