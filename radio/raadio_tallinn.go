package radio

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const raadioTallinnRDSURL = "https://raadiotallinn.err.ee/api/rds/getForChannel?channel=raadiotallinn"

func NewRaadioTallinn(httpClient *http.Client) *raadioTallinn {
	return &raadioTallinn{
		httpClient: httpClient,
	}
}

type raadioTallinn struct {
	httpClient *http.Client
}

func (r *raadioTallinn) CurrentTrack() (string, error) {
	res, err := r.httpClient.Get(raadioTallinnRDSURL)
	if err != nil {
		return "", fmt.Errorf("RDS request: %w", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("read RDS response: %w", err)
	}

	var rdsResponse struct {
		RDS string `json:"rds"`
	}

	if err := json.Unmarshal(body, &rdsResponse); err != nil {
		return "", fmt.Errorf("unmarshal RDS response: %w", err)
	}

	return rdsResponse.RDS, nil
}
