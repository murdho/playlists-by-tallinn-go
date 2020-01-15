package playlistsbytallinn

import (
	"context"
	"fmt"

	"github.com/murdho/playlists-by-tallinn/track"
)

//go:generate moq -out machinery_moq_test.go . Radio TrackStorage Logger

type Radio interface {
	CurrentTrack() (string, error)
}

type TrackStorage interface {
	Load(ctx context.Context, name string) (*track.Track, error)
	Save(ctx context.Context, trk track.Track) error
}

type Logger interface {
	Infow(msg string, keysAndValues ...interface{})
	Debug(args ...interface{})
}

type Machinery struct {
	Radio        Radio
	TrackStorage TrackStorage
	Logger       Logger
}

func NewMachinery(opts ...MachineryOption) Machinery {
	m := Machinery{}

	for _, opt := range opts {
		opt(&m)
	}

	return m
}

func (m *Machinery) Run(ctx context.Context) error {
	m.Logger.Debug("running")

	trackName, err := m.Radio.CurrentTrack()
	if err != nil {
		return fmt.Errorf("get current track: %w", err)
	}

	m.Logger.Infow("current track", "name", trackName)

	if trackName == "" {
		m.Logger.Debug("no current track, all done")
		return nil
	}

	if err := m.StoreTrack(ctx, trackName); err != nil {
		return fmt.Errorf("store track: %w", err)
	}

	return nil
}

func (m *Machinery) StoreTrack(ctx context.Context, name string) error {
	m.Logger.Debug("loading track from storage")

	trk, err := m.TrackStorage.Load(ctx, name)
	if err != nil {
		return fmt.Errorf("load track from storage: %w", err)
	}

	if trk != nil {
		m.Logger.Debug("track is already in storage")
		return nil
	}

	m.Logger.Debug("saving track to storage")

	if err := m.TrackStorage.Save(ctx, track.New(name)); err != nil {
		return fmt.Errorf("save track to storage: %w", err)
	}

	m.Logger.Debug("track saved to storage")
	return nil
}

type MachineryOption func(m *Machinery)

func WithRadio(r Radio) MachineryOption {
	return func(m *Machinery) {
		m.Radio = r
	}
}

func WithTrackStorage(ts TrackStorage) MachineryOption {
	return func(m *Machinery) {
		m.TrackStorage = ts
	}
}

func WithLogger(l Logger) MachineryOption {
	return func(m *Machinery) {
		m.Logger = l
	}
}
