package internal

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"log"
)

func NewLogger(debug bool) *zap.Logger {
	config := zap.NewProductionConfig()

	if debug {
		config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}

	logger, err := config.Build()
	if err != nil {
		log.Fatal(errors.Wrap(err, "building logger failed"))
	}

	return logger
}
