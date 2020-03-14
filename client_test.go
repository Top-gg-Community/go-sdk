package dbl

import (
	"log"
	"net/http"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	const clientTimeout = 5 * time.Second

	httpClient := &http.Client{}

	_, err := NewClient(
		"token",
		HTTPClientOption(httpClient), // Setting a custom HTTP client. Default is *http.Client with default timeout.
		TimeoutOption(clientTimeout), // Setting timeout option. Default is 3 seconds
	)
	if err != nil {
		log.Fatalf("Error creating new Discord Bot List client: %s", err)
	}

	// ...
}
