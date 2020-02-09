package playlistsbytallinn

import (
	"os"
	"testing"
)

func TestEnvOrDefault(t *testing.T) {
	if err := os.Setenv("a", "b"); err != nil {
		t.Errorf("unexpected error: %+v", err)
	}

	if err := os.Unsetenv("c"); err != nil {
		t.Errorf("unexpected error: %+v", err)
	}

	b := envOrDefault("a", "x")
	if b != "b" {
		t.Errorf("b:\ngot  %+v\nwant %+v", b, "b")
	}

	y := envOrDefault("c", "y")
	if y != "y" {
		t.Errorf("c:\ngot  %+v\nwant %+v", y, "y")
	}
}
