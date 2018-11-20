package dbl

import (
	"net/url"
)

type LargeWidget struct {
	Widget
	data WidgetData
}

func NewLargeWidget() *LargeWidget {
	return &LargeWidget{
		data: WidgetData{
			values: url.Values(make(map[string][]string)),
		},
	}
}

func (w *LargeWidget) BotID(id string) *LargeWidget {
	w.data.botID = id

	return w
}

func (w *LargeWidget) Extension(extension Extension) *LargeWidget {
	w.data.extension = extension

	return w
}

func (w *LargeWidget) TopColor(color int64) *LargeWidget {
	w.data.setValue("topcolor", color)

	return w
}

func (w *LargeWidget) MiddleColor(color int64) *LargeWidget {
	w.data.setValue("middlecolor", color)

	return w
}

func (w *LargeWidget) UsernameColor(color int64) *LargeWidget {
	w.data.setValue("usernamecolor", color)

	return w
}

func (w *LargeWidget) CertifiedColor(color int64) *LargeWidget {
	w.data.setValue("certifiedcolor", color)

	return w
}

func (w *LargeWidget) DataColor(color int64) *LargeWidget {
	w.data.setValue("datacolor", color)

	return w
}

func (w *LargeWidget) LabelColor(color int64) *LargeWidget {
	w.data.setValue("labelcolor", color)

	return w
}

func (w *LargeWidget) HighlightColor(color int64) *LargeWidget {
	w.data.setValue("highlightcolor", color)

	return w
}

func (w *LargeWidget) Generate() string {
	u, _ := url.Parse(BaseURL + "widgets/" + w.data.botID + w.data.extension.Ext())

	u.RawQuery = w.data.values.Encode()

	return u.String()
}
