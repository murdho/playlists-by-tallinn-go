package playlistsbytallinn

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	rdsURL                  = "https://raadiotallinn.err.ee/api/rds/getForChannel?channel=raadiotallinn"
	firestoreCollectionName = "playlists-by-tallinn"
)

var (
	httpClient   *http.Client
	trackStorage TrackStorage
	gcpProject   string
)

func init() {
	gcpProject = os.Getenv("GCP_PROJECT")

	httpClient = &http.Client{
		Timeout: time.Second,
	}

	firestoreClient, err := firestore.NewClient(context.Background(), gcpProject)
	if err != nil {
		log.Fatal(errors.Wrap(err, "initialising Firestore client failed"))
	}

	trackStorage = NewFirestoreTrackHandler(firestoreClient, gcpProject, firestoreCollectionName)
}

func PlaylistsByTallinn(ctx context.Context, _ PubSubMessage) error {
	trackName, err := CurrentTrack()
	if err != nil {
		return errors.Wrap(err, "getting current track failed")
	}

	if trackName == "" {
		return nil
	}

	track, err := trackStorage.LoadTrack(ctx, trackName)
	if err != nil {
		return errors.Wrap(err, "loading track from storage failed")
	}

	defer func() { log.Println(track) }()

	if track.IsIgnored() || track.IsProcessed() {
		return nil
	}

	track.ProcessedAt = time.Now()

	if err := trackStorage.SaveTrack(ctx, track); err != nil {
		return errors.Wrap(err, "saving track to storage failed")
	}

	return nil
}

type PubSubMessage struct {
	Data []byte `json:"data"`
}
