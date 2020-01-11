package playlistsbytallinn

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	"github.com/murdho/playlists-by-tallinn/track"
)

type Radio interface {
	CurrentTrack() (string, error)
}

type TrackStorage interface {
	Load(ctx context.Context, name string) (*track.Track, error)
	Save(ctx context.Context, track track.Track) error
}

type Machinery struct {
	Radio        Radio
	TrackStorage TrackStorage
	Logger       *zap.SugaredLogger
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
		return fmt.Errorf("load track from storage: %w")
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

func WithLogger(l *zap.SugaredLogger) MachineryOption {
	return func(m *Machinery) {
		m.Logger = l
	}
}

func WithTrackStorage(ts TrackStorage) MachineryOption {
	return func(m *Machinery) {
		m.TrackStorage = ts
	}
}
