package playlistsbytallinn

import (
	"context"

	"go.uber.org/zap"

	"github.com/murdho/playlists-by-tallinn/internal"
)

var sys *system

func InitSystem(r Radio, ts TrackStorage, l *zap.Logger) {
	sys = &system{
		radio:        r,
		trackStorage: ts,
		logger:       l,
	}
}

type system struct {
	radio        Radio
	trackStorage TrackStorage
	logger       *zap.Logger
}

type Radio interface {
	CurrentTrack() (string, error)
}

type TrackStorage interface {
	LoadTrack(ctx context.Context, trackName string) (*internal.Track, error)
	SaveTrack(ctx context.Context, track *internal.Track) error
}
