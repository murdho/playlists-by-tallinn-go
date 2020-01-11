package playlistsbytallinn

import (
	"go.uber.org/zap"
)

type Option func(m *Machinery)

func WithRadio(r Radio) Option {
	return func(m *Machinery) {
		m.Radio = r
	}
}

func WithLogger(l *zap.SugaredLogger) Option {
	return func(m *Machinery) {
		m.Logger = l
	}
}

func WithTrackStorage(ts TrackStorage) Option {
	return func(m *Machinery) {
		m.TrackStorage = ts
	}
}
