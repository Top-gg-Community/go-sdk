package dbl

import (
	"io/ioutil"
	"net/http"
)

func readBody(res *http.Response) ([]byte, error) {
	defer res.Body.Close()

	if res.StatusCode == 401 {
		return nil, ErrUnauthorizedRequest
	}

	if res.StatusCode == 429 {
		return nil, ErrRemoteRatelimit
	}

	if res.StatusCode != 200 {
		return nil, ErrRequestFailed
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}
