package dbl

import (
	"net/url"
	"strconv"
)

type Extension int

const (
	SVG Extension = iota
	PNG
)

type Widget interface {
	BotID() Widget
	Generate() string
	Extension() Widget
}

type WidgetData struct {
	botID     string
	values    url.Values
	extension Extension
}

func (w *WidgetData) setValue(key string, value int64) {
	w.values.Add(key, strconv.FormatInt(value, 16))
}

func (e Extension) Ext() string {
	switch e {
	case SVG:
		{
			return ".svg"
		}
	case PNG:
		{
			return ".png"
		}
	default:
		{
			return ".svg"
		}
	}
}
