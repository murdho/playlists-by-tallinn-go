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
	debugLogger, err := logger.New(logger.WithLevel("debug"))
	if err != nil {
		log.Fatal(err)
	}

	m := pbt.NewMachinery(
		pbt.WithRadio(new(testRadio)),
		pbt.WithTrackStorage(new(testStorage)),
		pbt.WithLogger(debugLogger),
	)

	if err := m.Run(context.Background()); err != nil {
		log.Fatal(err)
	}
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

	track := track.New(currentTrack)
	return &track, nil
}

func (ts *testStorage) Save(ctx context.Context, track track.Track) error {
	return nil
}
