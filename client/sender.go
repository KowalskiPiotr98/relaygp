package client

import (
	mail "github.com/wneessen/go-mail"
)

type MessageOpts struct {
}

func SendMessage(opts MessageOpts) error {
	message := mail.NewMsg()

	if err := message.From("test@localhost"); err != nil {
		return err
	}
	if err := message.To("test@localhost"); err != nil {
		return err
	}
	message.Subject("test")
	message.SetBodyString(mail.TypeTextPlain, "fuck off")

	client, err := mail.NewClient(
		"localhost",
		mail.WithPort(2525),
		mail.WithTLSPolicy(mail.NoTLS),
	)
	if err != nil {
		return nil
	}
	return client.DialAndSend(message)
}
