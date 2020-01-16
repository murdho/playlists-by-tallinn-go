package playlistsbytallinn

import (
	"context"
	"reflect"
	"testing"

	"github.com/murdho/playlists-by-tallinn/track"
)

func TestNewMachinery(t *testing.T) {
	r := radioMock()
	ts := trackStorageMock()
	l := loggerMock()

	m := NewMachinery(
		WithRadio(r),
		WithTrackStorage(ts),
		WithLogger(l),
	)

	if m.Radio != r {
		t.Errorf("new machinery radio:\ngot  %+v\nwant %+v", m.Radio, r)
	}

	if m.TrackStorage != ts {
		t.Errorf("new machinery track storage:\ngot  %+v\nwant %+v", m.TrackStorage, ts)
	}

	if m.Logger != l {
		t.Errorf("new machinery logger:\ngot  %+v\nwant %+v", m.Logger, l)
	}
}

func TestMachinery_RunNewTrack(t *testing.T) {
	expectedTrack := track.Track{Name: "a"}

	r := radioMock()
	ts := trackStorageMock()
	l := loggerMock()

	r.CurrentTrackFunc = func() (string, error) {
		return "a", nil
	}

	ts.LoadFunc = func(ctx context.Context, name string) (*track.Track, error) {
		return nil, nil
	}

	var actualTrack track.Track
	ts.SaveFunc = func(ctx context.Context, trk track.Track) error {
		actualTrack = trk
		return nil
	}

	m := NewMachinery(
		WithRadio(r),
		WithTrackStorage(ts),
		WithLogger(l),
	)

	err := m.Run(context.Background())
	if err != nil {
		t.Errorf("unexpected error: %+v", err, )
	}

	if !reflect.DeepEqual(actualTrack, expectedTrack) {
		t.Errorf("new track:\ngot  %+v\nwant %+v", actualTrack, expectedTrack)
	}
}

func TestMachinery_RunExistingTrack(t *testing.T) {
	r := radioMock()
	ts := trackStorageMock()
	l := loggerMock()

	r.CurrentTrackFunc = func() (string, error) {
		return "a", nil
	}

	ts.LoadFunc = func(ctx context.Context, name string) (*track.Track, error) {
		trk := track.New("a")
		return &trk, nil
	}

	m := NewMachinery(
		WithRadio(r),
		WithTrackStorage(ts),
		WithLogger(l),
	)

	err := m.Run(context.Background())
	if err != nil {
		t.Errorf("unexpected error: %+v", err, )
	}

	if len(ts.SaveCalls()) != 0 {
		t.Errorf("track storage save:\ngot  %+v\nwant %+v", len(ts.SaveCalls()), 0)
	}
}

func TestMachinery_RunNoCurrentTrack(t *testing.T) {
	r := radioMock()
	ts := trackStorageMock()
	l := loggerMock()

	r.CurrentTrackFunc = func() (string, error) {
		return "", nil
	}

	m := NewMachinery(
		WithRadio(r),
		WithTrackStorage(ts),
		WithLogger(l),
	)

	err := m.Run(context.Background())
	if err != nil {
		t.Errorf("unexpected error: %+v", err, )
	}

	if len(ts.LoadCalls()) != 0 {
		t.Errorf("track storage load:\ngot  %+v\nwant %+v", len(ts.LoadCalls()), 0)
	}

	if len(ts.SaveCalls()) != 0 {
		t.Errorf("track storage save:\ngot  %+v\nwant %+v", len(ts.SaveCalls()), 0)
	}
}

func radioMock() *RadioMock {
	return &RadioMock{
		CurrentTrackFunc: func() (s string, err error) {
			return "", nil
		},
	}
}

func trackStorageMock() *TrackStorageMock {
	return &TrackStorageMock{
		LoadFunc: func(ctx context.Context, name string) (track *track.Track, err error) {
			return nil, nil
		},
		SaveFunc: func(ctx context.Context, trk track.Track) error {
			return nil
		},
	}
}

func loggerMock() *LoggerMock {
	return &LoggerMock{
		DebugFunc: func(args ...interface{}) {
		},
		InfowFunc: func(msg string, keysAndValues ...interface{}) {
		},
	}
}
