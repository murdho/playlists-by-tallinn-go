package storage

import (
	"context"
	"crypto/md5"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	track "github.com/murdho/playlists-by-tallinn/track"
)

func NewTrack(firestore Firestore) *trackStorage {
	return &trackStorage{
		firestore: firestore,
	}
}

type trackStorage struct {
	firestore Firestore
}

func (t *trackStorage) Load(ctx context.Context, name string) (*track.Track, error) {
	var trk track.Track
	if err := t.firestore.Get(ctx, &trk, documentID(name)); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, nil
		}

		return nil, fmt.Errorf("get track from firestore: %w", err)
	}

	return &trk, nil
}

func (t *trackStorage) Save(ctx context.Context, track track.Track) error {
	if err := t.firestore.Set(ctx, documentID(track.Name), track); err != nil {
		return fmt.Errorf("add track to firestore: %w", err)
	}

	return nil
}

func documentID(name string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(name)))
}
