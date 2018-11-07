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

	assert.Nil(t, err)

	assert.Equal(t, len(bots.Results), fetchLimit, "Results array size should match request limit")
	assert.Equal(t, bots.Total, fetchLimit, "Total count should match request limit")
	assert.Equal(t, bots.Offset, 0, "Offset should be zero or non-specified")
}

func TestBot(t *testing.T) {
	client, err := NewClient("Unauthenticated request")

	if err != nil {
		t.Error("Client creation error", err)
	}

	bot, err := client.GetBot(testBotID)

	assert.Nil(t, err, "Unable to get user data")

	assert.Equal(t, bot.ID, testBotID, "Request & result bot ID should match")
}
