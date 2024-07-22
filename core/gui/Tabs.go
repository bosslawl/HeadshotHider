package gui

import (
	"HeadshotHider/core/backend/hide"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func Create(app fyne.App, window fyne.Window) *container.AppTabs {
	client := hide.NewClient(app)

	return &container.AppTabs{Items: []*container.TabItem{
		Dashboard(app, window, client).TabItem(),
		Settings(app, window, client).TabItem(),
		Credits(app).TabItem(),
	}}
}
