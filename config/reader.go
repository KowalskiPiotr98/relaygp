package config

import (
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

const VAR_PREFIX = "relaygp"

var (
	currentConfig config
)

type config struct {
	DebugMode bool
}

func init() {
	currentConfig = config{
		DebugMode: readOptionalEnvVar("debug", "false") == "true",
	}
}

func GetCurrentConfig() config {
	return currentConfig
}

func readRequiredVar(name string) string {
	value, ok := readEnvVar(name)
	if !ok {
		log.Panicf("Required env var %s not set", name)
	}
	return value
}

func readOptionalEnvVar(name string, fallback string) string {
	if value, ok := readEnvVar(name); ok {
		return value
	}
	return fallback
}

func readEnvVar(name string) (string, bool) {
	varName := strings.ToUpper(fmt.Sprintf("%s_%s", VAR_PREFIX, name))
	return os.LookupEnv(varName)
}
