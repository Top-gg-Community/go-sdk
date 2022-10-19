package dbl

import (
	"net/url"
	"strconv"
)

type DBL struct {
	Key string
}

type Extension int

const (
	SVG Extension = iota
	PNG
)

type Widget interface {
	Generate() string
	Extension() Widget
}

type WidgetData struct {
	botID     int
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

func (c *DBL) Widget(botID int, extension Extension) Widget{
	return &WidgetData{
		botID:     botID,
		values:    url.Values{},
		extension: extension,
	}
}

func (w *WidgetData) BotID() int {
	return w.botID
}

func (w *WidgetData) Extension() Extension {
	return w.extension
}

func (w *WidgetData) SetValue(key string, value int64) {
	w.values.Add(key, strconv.FormatInt(value, 16))
}
