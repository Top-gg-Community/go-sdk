package dbl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWeekend(t *testing.T) {
	client, err := NewClient("Unauthenticated request")

	assert.Nil(t, err, "Client should be created w/o error")

	_, err = client.IsMultiplierActive()

	assert.Nil(t, err, "Multiplier endpoint should not error (API depended)")
}
