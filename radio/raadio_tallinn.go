package radio

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func NewRaadioTallinn(opts ...option) raadioTallinn {
	rt := raadioTallinn{
		radio: new(radio),
	}

	for _, opt := range opts {
		opt(rt.radio)
	}

	return rt
}

type raadioTallinn struct {
	*radio
}

func (r raadioTallinn) CurrentTrack() (string, error) {
	res, err := r.httpClient.Get(r.url)
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
