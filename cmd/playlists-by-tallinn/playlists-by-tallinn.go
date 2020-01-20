package main

import (
	"context"
	"log"

	pbt "github.com/murdho/playlists-by-tallinn"
	"github.com/murdho/playlists-by-tallinn/logger"
	"github.com/murdho/playlists-by-tallinn/track"
)

const (
	// currentTrack = ""
	currentTrack         = "La La - Land Yo"
	trackExistsInStorage = false
	// trackExistsInStorage = true
)

func main() {
	// Firestore emulator:
	// 		gcloud beta emulators firestore start
	//
	// os.Setenv("GCP_PROJECT", "bla")
	// os.Setenv("FIRESTORE_EMULATOR_HOST", "8277")
	//
	// if err := runReal(); err != nil {
	// 	log.Fatal(err)
	// }

	if err := runFake(); err != nil {
		log.Fatal(err)
	}
}

func runReal() error {
	if err := pbt.PlaylistsByTallinn(context.Background(), struct{}{}); err != nil {
		return err
	}

	return nil
}

func runFake() error {
	debugLogger, err := logger.New(logger.WithLevel("debug"))
	if err != nil {
		return err
	}

	m := pbt.NewMachinery(
		pbt.WithRadio(new(testRadio)),
		pbt.WithTrackStorage(new(testStorage)),
		pbt.WithLogger(debugLogger),
	)

	if err := m.Run(context.Background()); err != nil {
		return err
	}

	return nil
}

type testRadio struct{}

func (tr *testRadio) CurrentTrack() (string, error) {
	return currentTrack, nil
}

type testStorage struct{}

func (ts *testStorage) Load(ctx context.Context, name string) (*track.Track, error) {
	if !trackExistsInStorage {
		return nil, nil
	}

	trk := track.New(currentTrack)
	return &trk, nil
}

func (ts *testStorage) Save(ctx context.Context, track track.Track) error {
	return nil
}
