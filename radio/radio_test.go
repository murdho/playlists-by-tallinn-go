package radio

import (
	"net/http"
	"testing"
)

func TestWithURL(t *testing.T) {
	r := &radio{}
	url := "a"

	WithURL(url)(r)

	if r.url != url {
		t.Errorf("radio URL:\ngot  %+v\nwant %+v", r.url, url)
	}
}

func TestWithHTTPClient(t *testing.T) {
	r := &radio{}
	hc := &http.Client{}

	WithHTTPClient(hc)(r)

	if r.httpClient != hc {
		t.Errorf("radio URL:\ngot  %+v\nwant %+v", r.httpClient, hc)
	}
}
