package gui

import (
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type credits struct {
	icon        *clickableIcon
	nameLabel   *widget.Label
	spacerLabel *widget.Label
	hyperlink   *widget.Hyperlink
	app         fyne.App
}

func Credits(a fyne.App) *credits {
	return &credits{app: a}
}

func (c *credits) BuildUI() *fyne.Container {
	const (
		https  = "https"
		github = "github.com"
	)

	repoURL := &url.URL{Scheme: https, Host: github, Path: "/bosslawl/HeadshotHider"}
	c.icon = CreateIcon(c.app.Icon(), repoURL, c.app)

	c.nameLabel = newBoldLabel("HeadshotHider")
	c.spacerLabel = newBoldLabel(" ")

	release := &url.URL{
		Scheme: https,
		Host:   github,
		Path:   "/bosslawl/HeadshotHider/releases/latest",
	}
	c.hyperlink = &widget.Hyperlink{Text: "Latest Release", URL: release, TextStyle: fyne.TextStyle{Bold: true}}

	spacer := &layout.Spacer{}
	return container.NewVBox(
		spacer,
		container.NewHBox(spacer, c.icon, spacer),
		container.NewHBox(spacer, c.nameLabel, spacer, c.hyperlink, spacer),
		container.NewHBox(spacer, newBoldLabel("Created with <3 by @bosslawl"), spacer),
		container.NewHBox(spacer, newBoldLabel("Special thanks to: @q3y, @deis., @onefivefiveseven"), spacer),
		spacer,
	)
}

func (c *credits) TabItem() *container.TabItem {
	return &container.TabItem{
		Text:    "Credits",
		Icon:    theme.InfoIcon(),
		Content: c.BuildUI(),
	}
}

type clickableIcon struct {
	widget.BaseWidget
	app  fyne.App
	url  *url.URL
	icon *canvas.Image
}

func (c *clickableIcon) Tapped(_ *fyne.PointEvent) {
	c.app.OpenURL(c.url)
}

func (c *clickableIcon) CreateRenderer() fyne.WidgetRenderer {
	c.ExtendBaseWidget(c)
	return widget.NewSimpleRenderer(c.icon)
}

func (c *clickableIcon) MinSize() fyne.Size {
	return fyne.Size{Width: 256, Height: 256}
}

func (c *clickableIcon) Cursor() desktop.Cursor {
	return desktop.PointerCursor
}

func CreateIcon(r fyne.Resource, u *url.URL, a fyne.App) *clickableIcon {
	return &clickableIcon{app: a, url: u, icon: canvas.NewImageFromResource(r)}
}
