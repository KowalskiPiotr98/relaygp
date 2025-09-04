package client

import (
	"net/smtp"
)

type MessageOpts struct {
	To   string
	Data []byte
}

func SendMessage(opts MessageOpts) error {
	return smtp.SendMail(
		"localhost:2525",
		nil,
		"test@localhost",
		[]string{opts.To},
		opts.Data,
	)
}
