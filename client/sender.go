package client

import (
	"net/smtp"
	"relaygp/config"
)

type MessageOpts struct {
	To   string
	Data []byte
}

func SendMessage(opts MessageOpts) error {
	config := config.GetCurrentConfig().SenderConfig

	var auth smtp.Auth

	if config.SmtpUser != "" && config.SmtpPassword != "" {
		auth = smtp.PlainAuth("", config.SmtpUser, config.SmtpPassword, config.SmtpAddress)
	}

	return smtp.SendMail(
		config.SmtpAddress,
		auth,
		config.SenderName,
		[]string{opts.To},
		opts.Data,
	)
}
