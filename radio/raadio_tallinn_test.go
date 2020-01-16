package radio

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewRaadioTallinn(t *testing.T) {
	url := "a"
	httpClient := &http.Client{}

	rt := NewRaadioTallinn(
		WithURL(url),
		WithHTTPClient(httpClient),
	)

	if rt.url != url {
		t.Errorf("raadio tallinn url:\ngot  %+v\nwant %+v", rt.url, url)
	}

	if rt.httpClient != httpClient {
		t.Errorf("raadio tallinn http client:\ngot  %+v\nwant %+v", rt.httpClient, httpClient)
	}
}

func TestRaadioTallinn_CurrentTrack(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if _, err := w.Write([]byte(`{"rds":"a"}`)); err != nil {
			t.Error(err)
		}
	}))
	defer ts.Close()

	httpClient := &http.Client{}
	rt := NewRaadioTallinn(
		WithURL(ts.URL),
		WithHTTPClient(httpClient),
	)

	currentTrack, err := rt.CurrentTrack()
	if err != nil {
		t.Errorf("unexpected error: %+v", err)
	}

	if currentTrack != "a" {
		t.Errorf("current track:\ngot  %+v\nwant %+v", currentTrack, "a")
	}
}
