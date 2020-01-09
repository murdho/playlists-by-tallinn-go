package lazyhttp

import (
	"net/http"
	"sync"
	"time"
)

var once *sync.Once
var client *http.Client

func Client() *http.Client {
	once.Do(func() {
		client = &http.Client{Timeout: 2 * time.Second}
	})

	return client
}
