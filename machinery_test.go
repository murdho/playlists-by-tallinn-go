package playlistsbytallinn

import (
	"testing"
)

func TestNewMachinery(t *testing.T) {
	r := new(RadioMock)
	ts := new(TrackStorageMock)
	l := new(LoggerMock)

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


