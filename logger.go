package main

import (
	"errors"
	"fmt"
	"github.com/op/go-logging"
	"os"
)

var Logger *logging.Logger

func init() {
	logLevelString := os.Getenv("LOG_LEVEL")
	logLevel, levelErr := parseLogLevel(logLevelString)
	logging.SetLevel(logLevel, "full")
	var err error
	Logger, err = logging.GetLogger("full")
	if err != nil {
		fmt.Printf("cannot get logger")
		panic(err)
	}
	if levelErr != nil {
		Logger.Error(levelErr.Error())
	}
}

func parseLogLevel(levelString string) (logging.Level, error) {
	switch levelString {
	case "CRITICAL":
		return logging.CRITICAL, nil
	case "ERROR":
		return logging.ERROR, nil
	case "WARNING":
		return logging.WARNING, nil
	case "NOTICE":
		return logging.NOTICE, nil
	case "INFO":
		return logging.INFO, nil
	case "DEBUG":
		return logging.DEBUG, nil

	default:
		return logging.INFO, errors.New("cannot parse log level for golang-image-processing." +
			" log level is set as INFO")
	}
}
