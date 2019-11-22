package logger

import (
	"log"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Level int

const (
	DebugLevel Level = iota
	InfoLevel
)

func New(level Level) *zap.Logger {
	config := zap.NewProductionConfig()

	if level == DebugLevel {
		config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}

	logger, err := config.Build()
	if err != nil {
		log.Fatal(errors.Wrap(err, "building logger failed"))
	}

	return logger
}
