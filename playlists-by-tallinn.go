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

const raadioTallinnURL = "https://raadiotallinn.err.ee/api/rds/getForChannel?channel=raadiotallinn"

func PlaylistsByTallinn(ctx context.Context, _ struct{}) error {
	gcpProject, ok := os.LookupEnv("GCP_PROJECT")
	if !ok {
		return fmt.Errorf("environment variable GCP_PROJECT required")
	}

	firestoreCollection := envOrDefault("FIRESTORE_COLLECTION", "playlists-by-tallinn")
	logLevel := envOrDefault("LOG_LEVEL", "info")

	log, err := logger.New(logger.WithLevel(logLevel))
	if err != nil {
		return fmt.Errorf("new logger: %w", err)
	}

	httpClient := &http.Client{Timeout: 2 * time.Second}
	raadioTallinn := radio.NewRaadioTallinn(
		radio.WithURL(raadioTallinnURL),
		radio.WithHTTPClient(httpClient),
	)

	firestoreClient, err := firestore.New(
		ctx,
		firestore.Project(gcpProject),
		firestore.Collection(firestoreCollection),
	)
	if err != nil {
		return fmt.Errorf("new firestore client: %w", err)
	}
	defer func() {
		if err := firestoreClient.Close(); err != nil {
			log.Error(err)
		}
	}()

	trackStorage := storage.NewTrack(firestoreClient)

	m := NewMachinery(
		WithRadio(raadioTallinn),
		WithTrackStorage(trackStorage),
		WithLogger(log),
	)

	if err := m.Run(ctx); err != nil {
		return err
	}

	return nil
}

func envOrDefault(key, fallback string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	return val
}
