package playlistsbytallinn

import (
	"context"
	"os"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/murdho/playlists-by-tallinn/internal"
	"github.com/murdho/playlists-by-tallinn/internal/logger"
	"github.com/murdho/playlists-by-tallinn/radio"
	"github.com/murdho/playlists-by-tallinn/storage"
)

func PlaylistsByTallinn(ctx context.Context, _ struct{}) error {
	gcpProject, ok := os.LookupEnv("GCP_PROJECT")
	if !ok {
		return errors.New("environment variable GCP_PROJECT undefined")
	}

	logLevel := logger.InfoLevel
	if os.Getenv("DEBUG") != "" {
		logLevel = logger.DebugLevel
	}

	raadioTallinn := radio.NewRaadioTallinn()
	trackStorage := storage.NewFirestoreStorage(gcpProject, "playlists-by-tallinn")
	log := logger.New(logLevel)

	err := Run(
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

func Run(ctx context.Context, opts ...option) error {
	machinery := &Machinery{}

	for _, opt := range opts {
		opt(machinery)
	}

	if err := machinery.Run(ctx); err != nil {
		return err
	}

	return nil
}

type Machinery struct {
	Radio        Radio
	TrackStorage TrackStorage
	Logger       *zap.Logger
}

func (m *Machinery) Run(ctx context.Context) error {
	m.Logger.Debug("running")

	trackName, err := m.Radio.CurrentTrack()
	if err != nil {
		return errors.Wrap(err, "getting current track failed")
	}

	m.Logger.Info("current track", zap.String("name", trackName))

	if trackName == "" {
		m.Logger.Debug("current track empty, all done")
		return nil
	}

	if err := m.StoreTrack(ctx, trackName); err != nil {
		return err
	}

	return nil
}

func (m *Machinery) StoreTrack(ctx context.Context, trackName string) error {
	m.Logger.Debug("loading track from storage")

	track, err := m.TrackStorage.LoadTrack(ctx, trackName)
	if err != nil {
		return errors.Wrap(err, "loading track from storage failed")
	}

	m.Logger.Debug(
		"track from storage",
		zap.String("name", track.Name),
		zap.Bool("persists", track.Persists),
	)

	if track.Persists {
		m.Logger.Debug("track already persists, all done")
		return nil
	}

	track.Persists = true
	m.Logger.Debug("saving track to storage")

	if err := m.TrackStorage.SaveTrack(ctx, track); err != nil {
		return errors.Wrap(err, "saving track to storage failed")
	}

	m.Logger.Debug("track saved to storage, all done")
	return nil
}

type Radio interface {
	CurrentTrack() (string, error)
}

type TrackStorage interface {
	LoadTrack(ctx context.Context, trackName string) (*internal.Track, error)
	SaveTrack(ctx context.Context, track *internal.Track) error
}

type option func(m *Machinery)

func WithRadio(r Radio) option {
	return func(m *Machinery) {
		m.Radio = r
	}
}

func WithLogger(l *zap.Logger) option {
	return func(m *Machinery) {
		m.Logger = l
	}
}

func WithTrackStorage(ts TrackStorage) option {
	return func(m *Machinery) {
		m.TrackStorage = ts
	}
}
