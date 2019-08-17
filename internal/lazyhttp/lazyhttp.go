package lazyhttp

import (
	"net/http"
	"time"
)

var client *http.Client

func NewClient() *http.Client {
	if client == nil {
		client = &http.Client{Timeout: time.Second}
	}

	return client
}
