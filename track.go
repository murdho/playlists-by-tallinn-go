package playlistsbytallinn

import (
	"context"
	"fmt"
	"time"
)

type Track struct {
	Name        string    `firestore:"name"`
	Ignored     bool      `firestore:"ignored"`
	ProcessedAt time.Time `firestore:"processed_at"`
}

type TrackStorage interface {
	LoadTrack(ctx context.Context, trackName string) (*Track, error)
	SaveTrack(ctx context.Context, track *Track) error
}

func (track *Track) String() string {
	return fmt.Sprintf(
		"track: '%s', ignored: %t, processed_at: %s",
		track.Name,
		track.Ignored,
		track.ProcessedAt.Format("2006-01-02 15:04:05 -0700"),
	)
}

func (track *Track) IsIgnored() bool {
	return track.Ignored
}

func (track *Track) IsProcessed() bool {
	return !track.ProcessedAt.IsZero()
}
