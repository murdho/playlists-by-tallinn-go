package playlistsbytallinn

import (
	"context"
	"github.com/murdho/playlists-by-tallinn/radio"
	"github.com/murdho/playlists-by-tallinn/storage"
	"github.com/pkg/errors"
	"log"
	"os"
)

var System *system

func NewSystem(raadioTallinn radio.Radio, trackStorage storage.Client) *system {
	return &system{
		raadioTallinn: raadioTallinn,
		trackStorage:  trackStorage,
	}
}

type system struct {
	raadioTallinn radio.Radio
	trackStorage  storage.Client
}

func init() {
	gcpProject := os.Getenv("GCP_PROJECT")
	System = NewSystem(
		radio.NewRaadioTallinn(),
		storage.NewFirestoreClient(gcpProject, "playlists-by-tallinn"),
	)
}

func PlaylistsByTallinn(ctx context.Context, _ PubSubMessage) error {
	trackName, err := System.raadioTallinn.CurrentTrack()
	if err != nil {
		return errors.Wrap(err, "getting current track failed")
	}

	if trackName == "" {
		return nil
	}

	track, err := System.trackStorage.LoadTrack(ctx, trackName)
	if err != nil {
		return errors.Wrap(err, "loading track from storage failed")
	}

	defer func() { log.Println(track) }()

	if track.Persists {
		return nil
	}

	if err := System.trackStorage.SaveTrack(ctx, track); err != nil {
		return errors.Wrap(err, "saving track to storage failed")
	}

	return nil
}

type PubSubMessage struct {
	Data []byte `json:"data"`
}
