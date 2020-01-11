package playlistsbytallinn

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/murdho/playlists-by-tallinn/firestore"
	"github.com/murdho/playlists-by-tallinn/logger"
	"github.com/murdho/playlists-by-tallinn/radio"
	"github.com/murdho/playlists-by-tallinn/storage"
)

func PlaylistsByTallinn(ctx context.Context, _ struct{}) error {
	gcpProject, ok := os.LookupEnv("GCP_PROJECT")
	if !ok {
		return fmt.Errorf("environment variable GCP_PROJECT required")
	}

	firestoreCollection, ok := os.LookupEnv("FIRESTORE_COLLECTION")
	if !ok {
		firestoreCollection = "playlists-by-tallinn"
	}

	logLevel, ok := os.LookupEnv("LOG_LEVEL")
	if !ok {
		logLevel = "info"
	}

	log, err := logger.New(logger.WithLevel(logLevel))
	if err != nil {
		return fmt.Errorf("new logger: %w", err)
	}

	httpClient := &http.Client{Timeout: 2 * time.Second}
	raadioTallinn := radio.NewRaadioTallinn(httpClient)

	firestoreClient, err := firestore.New(
		ctx,
		firestore.Project(gcpProject),
		firestore.Collection(firestoreCollection),
	)
	if err != nil {
		return fmt.Errorf("new firestore client: %w", err)
	}

	trackStorage := storage.NewTrack(firestoreClient)

	err = Run(
		ctx,
		WithRadio(raadioTallinn),
		WithTrackStorage(trackStorage),
		WithLogger(log),
	)

	if err != nil {
		return err
	}

	return nil
}

func Run(ctx context.Context, opts ...MachineryOption) error {
	machinery := &Machinery{}

	for _, opt := range opts {
		opt(machinery)
	}

	if err := machinery.Run(ctx); err != nil {
		return err
	}

	return nil
}
