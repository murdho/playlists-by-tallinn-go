package radio

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/pkg/errors"

	"github.com/murdho/playlists-by-tallinn/internal/lazyhttp"
)

const raadioTallinnRDSURL = "https://raadiotallinn.err.ee/api/rds/getForChannel?channel=raadiotallinn"

func NewRaadioTallinn() *raadioTallinn {
	return new(raadioTallinn)
}

type raadioTallinn struct{}

func (r *raadioTallinn) CurrentTrack() (string, error) {
	res, err := lazyhttp.NewClient().Get(raadioTallinnRDSURL)
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

	var responseBody raadioTallinnRDSResponse
	if err := json.Unmarshal(body, &responseBody); err != nil {
		return "", errors.Wrap(err, "unmarshalling RDS response body failed")
	}

	return responseBody.RDS, nil

}

type raadioTallinnRDSResponse struct {
	RDS string `json:"rds"`
}
