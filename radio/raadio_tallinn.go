package radio

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
)

const rdsURL = "https://raadiotallinn.err.ee/api/rds/getForChannel?channel=raadiotallinn"

func NewRaadioTallinn() Radio {
	return &raadioTallinn{
		rdsURL:         rdsURL,
		httpClientFunc: getLazyHTTPClient,
	}
}

type raadioTallinn struct {
	rdsURL         string
	httpClientFunc func() httpClient
}

func (r *raadioTallinn) CurrentTrack() (string, error) {
	res, err := r.httpClientFunc().Get(r.rdsURL)
	if err != nil {
		return "", errors.Wrap(err, "RDS request failed")
	}
	defer func() {
		if err := res.Body.Close(); err != nil {
			log.Fatal(errors.Wrap(err, "closing RDS response body failed"))
		}
	}()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", errors.Wrap(err, "reading RDS response body failed")
	}

	var responseBody rdsResponse
	if err := json.Unmarshal(body, &responseBody); err != nil {
		return "", errors.Wrap(err, "unmarshalling RDS response body failed")
	}

	return responseBody.RDS, nil

}

type rdsResponse struct {
	RDS string `json:"rds"`
}
