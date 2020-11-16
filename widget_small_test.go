package dbl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	expectedSmallWidget       = BaseURL + "widgets/servers/574652751745777665.svg?leftcolor=555798&righttextcolor=f12350"
	expectedSmallWidgetNoType = BaseURL + "widgets/status/574652751745777665.svg?avatarbg=987230&lefttextcolor=123890"
)

func TestWidgetSmall(t *testing.T) {
	w := NewSmallWidget().
		BotID(testBotID).
		WidgetType(WidgetServers).
		LeftColor(0x555798).
		RightTextColor(0xf12350).
		Generate()

	assert.Equal(t, expectedSmallWidget, w)
}

func TestWidgetSmallNoType(t *testing.T) {
	w := NewSmallWidget().
		BotID(testBotID).
		AvatarBackground(0x987230).
		LeftTextColor(0x123890).
		Generate()

	assert.Equal(t, expectedSmallWidgetNoType, w)
}
