package dbl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testUserID = "105122038586286080"
)

func TestUsers(t *testing.T) {
	client, err := NewClient("Unauthenticated request")

	if err != nil {
		t.Error("Client creation error", err)
	}

	user, err := client.GetUser(testUserID)

	assert.Nil(t, err, "Unable to get user data")

	assert.Equal(t, user.ID, testUserID, "Request & response user ID should match")
}
