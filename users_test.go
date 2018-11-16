package dbl

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testUserID = "105122038586286080"
)

func TestUsers(t *testing.T) {
	client, err := NewClient(os.Getenv("apikey"))

	assert.Nil(t, err, "Client should be created w/o error")

	user, err := client.GetUser(testUserID)

	assert.Nil(t, err, "Unable to get user data")

	assert.Equal(t, testUserID, user.ID, "Request & response user ID should match")
}
