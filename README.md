## Go DBL

[![Build Status](https://travis-ci.com/rumblefrog/go-dbl.svg?branch=master)](https://travis-ci.com/rumblefrog/go-dbl)
[![Go Report Card](https://goreportcard.com/badge/github.com/rumblefrog/go-dbl)](https://goreportcard.com/report/github.com/rumblefrog/go-dbl)
[![GoDoc](https://godoc.org/github.com/rumblefrog/go-dbl?status.svg)](https://godoc.org/github.com/rumblefrog/go-dbl)

An API wrapper for [Discord Bots](https://discordbots.org/)

Godoc is available here: https://godoc.org/github.com/rumblefrog/go-dbl

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**

- [Guides](#guides)
  - [Installing](#installing)
  - [Posting Stats](#posting-stats)
  - [Timeout option](#timeout-option)
  - [Ratelimits](#ratelimits)
  - [More details](#more-details)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Guides

### Installing

```bash
go get -u github.com/rumblefrog/go-dbl
```

### Posting Stats

```go
func main() {
	client, err := dbl.NewClient("Token")
	
	if err != nil {
		// Handle error
	}
	
	err = client.PostBotStats("441751906428256277", &BotStatsPayload{
		Shards: []int{2500}, // If non-sharded, just pass total server count as the only integer element
	})
	
	if err != nil {
		// Handle error
	}
	
	// ...
}
```

### Timeout option

```go
func main() {
	client, err := dbl.NewClient(
		"Token",
		dbl.TimeoutOption(time.Second * 5) // // Setting timeout option. Default is 3 seconds
	)
	
	if err != nil {
		// Handle error
	}
	// ...
}
```

### Ratelimits

There's a local token bucket rate limiter, allowing for 60 requests a minute (single/burst)

Upon reaching the local rate limit, `ErrLocalRatelimit` error will be returned

If remote rate limit is exceeded, `ErrRemoteRatelimit` error will be returned and `RetryAfter` in client fields will be updated with the retry time

### More details

For more details, Godoc and tests are available
