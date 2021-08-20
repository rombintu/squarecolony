package utils

import (
	log "github.com/sirupsen/logrus"
)

func NewLogger() log.Logger {
	logger := log.New()
	logger.SetLevel(log.DebugLevel)
	// log.SetFormatter(&log.JSONFormatter{})
	logger.SetFormatter(logger.Formatter)
	return *logger
}
