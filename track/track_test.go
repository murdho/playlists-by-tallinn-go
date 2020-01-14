package track

import (
	"testing"
)

func TestNew(t *testing.T) {
	trk := New("a")

	if trk.Name != "a" {
		t.Errorf("track name:\ngot  %+v\nwant: %+v", trk.Name, "a")
	}
}
