package dbl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	expectedLargeWidget      = BaseURL + "widgets/574652751745777665.png?certifiedcolor=343434&topcolor=fffff"
	expectedLargeWidgetNoExt = BaseURL + "widgets/574652751745777665.svg?datacolor=309839"
)

func TestWidgetLarge(t *testing.T) {
	w := NewLargeWidget().
		BotID(testBotID).
		Extension(PNG).
		TopColor(0xfffff).
		CertifiedColor(0x343434).
		Generate()

	assert.Equal(t, expectedLargeWidget, w)
}

func TestWidgetLargeNoExt(t *testing.T) {
	w := NewLargeWidget().
		BotID(testBotID).
		DataColor(0x309839).
		Generate()

	assert.Equal(t, expectedLargeWidgetNoExt, w)
}
