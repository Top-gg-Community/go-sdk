## Go DBL

[![Build Status](https://travis-ci.com/rumblefrog/go-dbl.svg?branch=master)](https://travis-ci.com/rumblefrog/go-dbl)
[![Go Report Card](https://goreportcard.com/badge/github.com/DiscordBotList/go-dbl)](https://goreportcard.com/report/github.com/DiscordBotList/go-dbl)
[![GoDoc](https://godoc.org/github.com/DiscordBotList/go-dbl?status.svg)](https://godoc.org/github.com/DiscordBotList/go-dbl)

An API wrapper for [Discord Bots](https://top.gg/)

Godoc is available here: https://godoc.org/github.com/DiscordBotList/go-dbl

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
## Table of Contents

- [Guides](#guides)
  - [Installing](#installing)
  - [Posting Stats](#posting-stats)
  - [Setting options](#setting-options)
  - [Ratelimits](#ratelimits)
  - [Webhook](#webhook)
  - [More details](#more-details)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Guides

### Installing

```bash
go get -u github.com/DiscordBotList/go-dbl
```

### Posting Stats

```go
package main

import (
	"log"

	"github.com/DiscordBotList/go-dbl"
)

func main() {
	dblClient, err := dbl.NewClient("token")
	if err != nil {
		log.Fatalf("Error creating new Discord Bot List client: %s", err)
	}

	err = dblClient.PostBotStats("botID", &dbl.BotStatsPayload{
		Shards: []int{2500}, // If non-sharded, just pass total server count as the only integer element
	})
	if err != nil {
		log.Printf("Error sending bot stats to Discord Bot List: %s", err)
	}

	// ...
}
```

### Setting options

```go
package main

import (
	"log"
	"net/http"
	"time"

	"github.com/DiscordBotList/go-dbl"
)

const clientTimeout = 5 * time.Second

func main() {
	httpClient := &http.Client{}

	dblClient, err := dbl.NewClient(
		"token",
		dbl.HTTPClientOption(httpClient), // Setting a custom HTTP client. Default is *http.Client with default timeout.
		dbl.TimeoutOption(clientTimeout), // Setting timeout option. Default is 3 seconds
	)
	if err != nil {
		log.Fatalf("Error creating new Discord Bot List client: %s", err)
	}

	// ...
}
```

### Ratelimits

There's a local token bucket rate limiter, allowing for 60 requests a minute (single/burst)

Upon reaching the local rate limit, `ErrLocalRatelimit` error will be returned

If remote rate limit is exceeded, `ErrRemoteRatelimit` error will be returned and `RetryAfter` in client fields will be updated with the retry time

### Webhook

```go
package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/DiscordBotList/go-dbl"
)

const listenerPort = ":9090"

func main() {
	listener := dbl.NewListener("token", handleVote)

	// Serve is a blocking call
	err := listener.Serve(listenerPort)
	if !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("HTTP server error: %s", err)
	}
}

func handleVote(payload *dbl.WebhookPayload) {
	// perform on payload
}
```

### More details

For more details, Godoc and tests are available
