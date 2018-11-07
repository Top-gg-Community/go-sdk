package dbl

import (
	"encoding/json"
)

type weekendResponse struct {
	IsWeekend bool `json:"is_weekend"`
}

func (c *DBLClient) IsMultiplierActive() (bool, error) {
	res, err := c.client.Get(BaseURL + "weekend")

	if err != nil {
		return false, err
	}

	body, err := readBody(res)

	if err != nil {
		return false, err
	}

	wr := &weekendResponse{}

	err = json.Unmarshal(body, wr)

	if err != nil {
		return false, err
	}

	return wr.IsWeekend, nil
}
