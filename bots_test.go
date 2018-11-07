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

	assert.Nil(t, err, "Client should be created w/o error")

	bots, err := client.GetBots(&GetBotsPayload{
		Limit: fetchLimit,
	})

	assert.Nil(t, err, "Request should be successful (API depended)")

	assert.Equal(t, len(bots.Results), fetchLimit, "Results array size should match request limit")
	assert.Equal(t, bots.Count, fetchLimit, "Count should match request limit")
	assert.Equal(t, bots.Offset, 0, "Offset should be zero or non-specified")
}

func TestBot(t *testing.T) {
	client, err := NewClient("Unauthenticated request")

	assert.Nil(t, err, "Client should be created w/o error")

	bot, err := client.GetBot(testBotID)

	assert.Nil(t, err, "Unable to get user data")

	assert.Equal(t, bot.ID, testBotID, "Request & result bot ID should match")
}
