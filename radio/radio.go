package radio

import (
	"net/http"
	"time"
)

type Radio interface {
	CurrentTrack() (string, error)
}

var lazyHTTPClient httpClient

type httpClient interface {
	Get(URL string) (*http.Response, error)
}

func getLazyHTTPClient() httpClient {
	if lazyHTTPClient == nil {
		lazyHTTPClient = &http.Client{Timeout: time.Second}
	}

	return lazyHTTPClient
}
