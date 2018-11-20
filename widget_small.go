package dbl

import "net/url"

type WidgetType int

const (
	WidgetStatus WidgetType = iota
	WidgetUpvotes
	WidgetServers
	WidgetLib
)

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

type SmallWidget struct {
	Widget
	data  WidgetData
	wType WidgetType
}

func NewSmallWidget() *SmallWidget {
	return &SmallWidget{
		data: WidgetData{
			values: url.Values(make(map[string][]string)),
		},
	}
}

func (w *SmallWidget) BotID(id string) *SmallWidget {
	w.data.botID = id

	return w
}

func (w *SmallWidget) Extension(extension Extension) *SmallWidget {
	w.data.extension = extension

	return w
}

func (w *SmallWidget) WidgetType(wType WidgetType) *SmallWidget {
	w.wType = wType

	return w
}

func (w *SmallWidget) AvatarBackground(color int64) *SmallWidget {
	w.data.setValue("avatarbg", color)

	return w
}

func (w *SmallWidget) LeftColor(color int64) *SmallWidget {
	w.data.setValue("leftcolor", color)

	return w
}

func (w *SmallWidget) RightColor(color int64) *SmallWidget {
	w.data.setValue("rightcolor", color)

	return w
}

func (w *SmallWidget) LeftTextColor(color int64) *SmallWidget {
	w.data.setValue("lefttextcolor", color)

	return w
}

func (w *SmallWidget) RightTextColor(color int64) *SmallWidget {
	w.data.setValue("righttextcolor", color)

	return w
}

func (w *SmallWidget) Generate() string {
	u, _ := url.Parse(BaseURL + "widgets/" + w.wType.String() + "/" + w.data.botID + w.data.extension.Ext())

	u.RawQuery = w.data.values.Encode()

	return u.String()
}
