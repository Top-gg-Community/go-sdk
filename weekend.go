package dbl

import (
	"encoding/json"
)

type weekendResponse struct {
	IsWeekend bool `json:"is_weekend"`
}

// Check if the multiplier is live for the weekend
func (c *DBLClient) IsMultiplierActive() (bool, error) {
	req, err := c.createRequest("GET", "weekend", nil)

	if err != nil {
		return false, err
	}

	res, err := c.client.Do(req)

	if err != nil {
		return false, err
	}

	body, err := c.readBody(res)

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
