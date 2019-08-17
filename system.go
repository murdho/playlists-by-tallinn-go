package playlistsbytallinn

import (
	"github.com/murdho/playlists-by-tallinn/internal"
	"go.uber.org/zap"
)

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
	LoadTrack(trackName string) (*internal.Track, error)
	SaveTrack(*internal.Track) error
}
