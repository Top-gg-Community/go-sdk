package dbl

import "errors"

var (
	ErrRequestFailed         = errors.New("Remote request failed with non 200 status code")
	ErrLocalRatelimit        = errors.New("Exceeded local rate limit")
	ErrRemoteRatelimit       = errors.New("Exceeded remote rate limit")
	ErrUnauthorizedRequest   = errors.New("Unauthorized request")
	ErrRequireAuthentication = errors.New("Endpoint requires valid token")
)
