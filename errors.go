package dbl

import "errors"

var (
	ErrRequestFailed = errors.New("Remote request failed with non 200 status code")
)
