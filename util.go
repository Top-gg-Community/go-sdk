package dbl

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type ratelimitResponse struct {
	RetryAfter int `json:"retry-after"`
}

func (c *DBLClient) readBody(res *http.Response) ([]byte, error) {
	defer res.Body.Close()

	if res.StatusCode == 401 {
		return nil, ErrUnauthorizedRequest
	}

	if res.StatusCode != 200 {
		return nil, ErrRequestFailed
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	if res.StatusCode == 429 {
		rr := &ratelimitResponse{}

		err = json.Unmarshal(body, rr)

		if err != nil {
			return nil, err
		}

		// Not sure if mutex is needed for this one
		c.RetryAfter = rr.RetryAfter

		return nil, ErrRemoteRatelimit
	}

	return body, nil
}
