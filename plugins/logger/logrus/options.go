package logrus

import (
	"github.com/sirupsen/logrus"

	"cqsim-admin-core/logger"
)

type Options struct {
	logger.Options
	Formatter logrus.Formatter
	Hooks     logrus.LevelHooks
	// Flag for whether to log caller info (off by default)
	ReportCaller bool
	// Exit Function to call when FatalLevel log
	ExitFunc func(int)
}

type formatterKey struct{}

type hooksKey struct{}

type reportCallerKey struct{}

// warning to use this option. because logrus doest not open CallerDepth option
// this will only print this package

type exitKey struct{}

type logrusLoggerKey struct{}
