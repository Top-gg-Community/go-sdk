## Go DBL

[![Build Status](https://travis-ci.com/rumblefrog/go-dbl.svg?branch=master)](https://travis-ci.com/rumblefrog/go-dbl)
[![Go Report Card](https://goreportcard.com/badge/github.com/rumblefrog/go-dbl)](https://goreportcard.com/report/github.com/rumblefrog/go-dbl)
[![GoDoc](https://godoc.org/github.com/rumblefrog/go-dbl?status.svg)](https://godoc.org/github.com/rumblefrog/go-dbl)

An API wrapper for [Discord Bots](https://discordbots.org/)

Godoc is available here: https://godoc.org/github.com/rumblefrog/go-dbl


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

For more details, Godoc and tests are available
