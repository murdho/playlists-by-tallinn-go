package storage

import (
	"context"
)

type Client interface {
	LoadTrack(ctx context.Context, trackName string) (*Track, error)
	SaveTrack(ctx context.Context, track *Track) error
}

type Track struct {
	Name     string `firestore:"name"`
	Persists bool
}
