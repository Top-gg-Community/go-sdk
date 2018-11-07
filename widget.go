package dbl

import (
	"net/url"
)

type WidgetType int

const (
	WidgetStatus WidgetType = iota
	WidgetUpvotes
	WidgetServers
	WidgetLib
)

type Extension int

const (
	SVG Extension = iota
	PNG
)

type LargeWidgetOptions struct {
	TopColor       string
	MiddleColor    string
	UsernameColor  string
	CertifiedColor string
	DataColor      string
	LabelColor     string
	HighlightColor string
}

type SmallWidgetOptions struct {
	AvatarBackground string
	LeftColor        string
	RightColor       string
	LeftTextColor    string
	RightTextColor   string
	NoAvatar         *bool
}

func (t WidgetType) String() string {
	switch t {
	case WidgetStatus:
		{
			return "status"
		}
	case WidgetUpvotes:
		{
			return "upvotes"
		}
	case WidgetServers:
		{
			return "servers"
		}
	case WidgetLib:
		{
			return "lib"
		}
	default:
		{
			return "status"
		}
	}
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

func GetLargeWidget(botID string, options LargeWidgetOptions, ext Extension) string {
	u, _ := url.Parse(BaseURL + "widget/" + botID + ext.Ext())

	q := u.Query()

	// This is quite repetitive, there have to be a shorter way

	if options.TopColor != "" {
		q.Add("topcolor", options.TopColor)
	}

	if options.MiddleColor != "" {
		q.Add("middlecolor", options.MiddleColor)
	}

	if options.UsernameColor != "" {
		q.Add("usernamecolor", options.UsernameColor)
	}

	if options.CertifiedColor != "" {
		q.Add("certifiedcolor", options.CertifiedColor)
	}

	if options.DataColor != "" {
		q.Add("datacolor", options.DataColor)
	}

	if options.LabelColor != "" {
		q.Add("labelcolor", options.LabelColor)
	}

	if options.HighlightColor != "" {
		q.Add("highlightcolor", options.HighlightColor)
	}

	u.RawQuery = q.Encode()

	return u.String()
}

func GetSmallWidget(botID string, widgetType WidgetType, options SmallWidgetOptions, ext Extension) string {
	u, _ := url.Parse(BaseURL + "widget/" + widgetType.String() + "/" + botID + ext.Ext())

	q := u.Query()

	if options.AvatarBackground != "" {
		q.Add("avatarbg", options.AvatarBackground)
	}

	if options.LeftColor != "" {
		q.Add("leftcolor", options.LeftColor)
	}

	if options.RightColor != "" {
		q.Add("rightcolor", options.RightColor)
	}

	if options.LeftTextColor != "" {
		q.Add("lefttextcolor", options.LeftTextColor)
	}

	if options.RightTextColor != "" {
		q.Add("righttextcolor", options.RightTextColor)
	}

	u.RawQuery = q.Encode()

	return u.String()
}
