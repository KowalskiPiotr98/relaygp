package config

type clientConfig struct {
	SmtpAddress  string
	SmtpUser     string
	SmtpPassword string
	SenderName   string
}

func initClientConfig() clientConfig {
	return clientConfig{
		SmtpAddress:  readRequiredVar("sender_address"),
		SmtpUser:     readOptionalEnvVar("sender_user", ""),
		SmtpPassword: readOptionalEnvVar("sender_password", ""),
		SenderName:   readRequiredVar("sender_name"),
	}
}
