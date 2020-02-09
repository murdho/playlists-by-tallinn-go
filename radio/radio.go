package radio

import (
	"net/http"
)

type radio struct {
	url        string
	httpClient *http.Client
}

type option func(*radio)

func WithURL(url string) option {
	return func(rt *radio) {
		rt.url = url
	}
}

func WithHTTPClient(httpClient *http.Client) option {
	return func(rt *radio) {
		rt.httpClient = httpClient
	}
}
