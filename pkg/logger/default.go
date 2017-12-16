package logger

import (
	log "github.com/Sirupsen/logrus"
)

const (
	//LoggerHostname environment key
	LoggerHostname = "logger.hostname"
	//LoggerAppname environment key
	LoggerAppname = "logger.appname"
	//LoggerService environment key
	LoggerService = "logger.service"
	//LoggerVersion environment key
	LoggerVersion = "logger.version"
)

/**
//Logger global logger
var logger = log.WithFields(log.Fields{
	"hostname": envUtils.GetEnv(LoggerHostname, envUtils.GetEnv("macaque.hostname", "localhost")),
	"appname":  envUtils.GetEnv(LoggerAppname, envUtils.GetEnv("macaque.application", "macaque")),
	"service":  envUtils.GetEnv(LoggerService, envUtils.GetEnv("macaque.service", "")),
	"version":  envUtils.GetEnv(LoggerVersion, envUtils.GetEnv("macaque.version", "0.0.1-alpha1")),
})
**/
var logger = log.WithFields(log.Fields{})

//Info wrapper
func Info(args ...interface{}) {
	logger.Info(args)
}

//Infof wrapper
func Infof(format string, args ...interface{}) {
	logger.Infof(format, args)
}

//Debug wrapper
func Debug(args ...interface{}) {
	logger.Debug(args)
}

//Debugf wrapper
func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args)
}

//Warn wrapper
func Warn(args ...interface{}) {
	logger.Warn(args)
}

//Warnf wrapper
func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args)
}

//Error wrapper
func Error(args ...interface{}) {
	logger.Error(args)
}

//Errorf wrapper
func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args)
}
