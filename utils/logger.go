package utils

import (
	log "github.com/sirupsen/logrus"
)

func NewLogger(logLevel string) *log.Logger {
	logger := log.New()

	level, err := log.ParseLevel(logLevel)
	if err != nil {
		level = log.DebugLevel
	}

	logger.SetLevel(level)
	// log.SetFormatter(&log.JSONFormatter{})
	// log.SetFormatter(log.Formatter)
	return logger
}
