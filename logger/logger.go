package logger

import (
	"fmt"

	"go.uber.org/zap"
)

func New(opts ...Option) (*zap.SugaredLogger, error) {
	config := zap.NewProductionConfig()

	for _, opt := range opts {
		if err := opt(&config); err != nil {
			return nil, fmt.Errorf("apply opts: %w")
		}
	}

	logger, err := config.Build()
	if err != nil {
		return nil, fmt.Errorf("build logger: %w", err)
	}

	return logger.Sugar(), nil
}

type Option func(*zap.Config) error

// Possible levels:
// 	- debug
// 	- info
// 	- warn
// 	- error
// 	- dpanic
// 	- panic
// 	- fatal
func WithLevel(level string) Option {
	return func(c *zap.Config) error {
		l := zap.NewAtomicLevel()

		if err := l.UnmarshalText([]byte(level)); err != nil {
			return fmt.Errorf("unmarshal log level: %w", err)
		}

		c.Level = l
		return nil
	}
}
