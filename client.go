package dbl

import (
	"errors"
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

const (
	DefaultTimeout = time.Second * 3
)

var (
	ErrNilOption = errors.New("Invalid client option")
)

type DBLClient struct {
	// bots/* 60/m with 1 hour block if exceeded
	// Indicates how long the timeout period is/when you will be able to send requests again
	// Upon exceeding a rate limit, this will be updated with the retry-after value.
	RetryAfter int

	limiter *rate.Limiter

	client *http.Client

	timeout time.Duration

	token string
}

func TimeoutOption(duration time.Duration) func(*DBLClient) error {
	return func(c *DBLClient) error {
		c.client.Timeout = duration

		return nil
	}
}

func NewClient(token string, options ...func(*DBLClient) error) (c *DBLClient, err error) {
	c = &DBLClient{
		limiter: rate.NewLimiter(1, 60),
		client: &http.Client{
			Timeout: DefaultTimeout,
		},
		token: token,
	}

	for _, f := range options {
		if f == nil {
			return nil, ErrNilOption
		}
		if err = f(c); err != nil {
			return nil, err
		}
	}

	return c, nil
}
