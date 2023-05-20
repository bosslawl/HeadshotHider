// Package ui handles all logic related to the user interface.
package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/bosslawl/HeadshotHider/v2/internal/backend"
)

// Create will set up and create the ui components.
func Create(app fyne.App, window fyne.Window) *container.AppTabs {
	client := backend.NewClient(app)

	return &container.AppTabs{Items: []*container.TabItem{
		newManage(app, window, client).tabItem(),
		newSettings(app, window, client).tabItem(),
		newAbout(app).tabItem(),
	}}
}
