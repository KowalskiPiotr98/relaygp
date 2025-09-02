package config

import "strconv"

type listenerConfig struct {
	ListenAddress   string
	MaxMessageBytes int64
	MaxRecipients   int
}

func initListenerConfig() listenerConfig {
	maxRecipients := 1
	if recipients, err := strconv.Atoi(readOptionalEnvVar("max_recipients", "")); err != nil {
		maxRecipients = recipients
	}

	var maxMessageBytes int64 = 1024 * 1024
	if messageBytes, err := strconv.ParseInt(readOptionalEnvVar("max_message_bytes", ""), 10, 64); err != nil {
		maxMessageBytes = messageBytes
	}

	return listenerConfig{
		ListenAddress:   readOptionalEnvVar("listen", "localhost:2525"),
		MaxRecipients:   maxRecipients,
		MaxMessageBytes: maxMessageBytes,
	}
}

//todo: username + password
//todo: tls certificates
