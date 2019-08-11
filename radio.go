package playlistsbytallinn

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
)

type rdsResponse struct {
	RDS string `json:"rds"`
}

func CurrentTrack() (string, error) {
	res, err := httpClient.Get(rdsURL)
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
