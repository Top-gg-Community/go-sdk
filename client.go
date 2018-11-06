package dbl

import (
	"net/http"

	"golang.org/x/time/rate"
)

type DBLClient struct {
	limiter *rate.Limiter

	client *http.Client

	token string
}

func NewClient(token string) *DBLClient {
	return &DBLClient{
		limiter: rate.NewLimiter(1, 60),
		client:  &http.Client{},
		token:   token,
	}
}
