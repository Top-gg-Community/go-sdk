package dbl

type WidgetType int

const (
	WidgetOwner WidgetType = iota
	WidgetStatus
	WidgetUpvotes
	WidgetServers
	WidgetLib
)
