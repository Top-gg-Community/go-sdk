package dbl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testBotID  = "441751906428256277"
	fetchLimit = 20
)

func TestBots(t *testing.T) {
	client, err := NewClient("Unauthenticated request")

	if err != nil {
		t.Error("Client creation error", err)
	}

	bots, err := client.GetBots(&GetBotsPayload{
		Limit: fetchLimit,
	})

	assert.Equal(t, len(bots.Results), fetchLimit)
	assert.Equal(t, bots.Total, fetchLimit)
	assert.Equal(t, bots.Offset, 0)
}
