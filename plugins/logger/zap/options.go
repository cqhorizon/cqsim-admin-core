package zap

import (
	"github.com/cqhorizon/cqsim-admin-core/logger"
)

type Options struct {
	logger.Options
}

type callerSkipKey struct{}

func WithCallerSkip(i int) logger.Option {
	return logger.SetOption(callerSkipKey{}, i)
}

type configKey struct{}

// WithConfig pass zap.Config to logger

type encoderConfigKey struct{}

// WithEncoderConfig pass zapcore.EncoderConfig to logger

type namespaceKey struct{}

type writerKey struct{}
