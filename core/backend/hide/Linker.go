package hide

import (
	"fyne.io/fyne/v2"
)

type Client struct {
	app fyne.App

	DownloadPath string

	LoaderPath string
}

func NewClient(app fyne.App) *Client {
	return &Client{app: app}
}
