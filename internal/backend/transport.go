// Package transport handles sending and receiving using wormhole-william
package backend

import (
	"fyne.io/fyne/v2"
)

// Client defines the client for handling sending and receiving using wormhole-william
type Client struct {
	app fyne.App

	// DownloadPath holds the download path used for saving received files.
	DownloadPath string

	// LoaderPath holds the path to the loader used for sending files.
	LoaderPath string

	// Defines if we should pass a custom code or let wormhole-william generate on for us.
	CustomCode bool
}

// NewClient returns a new client for sending and receiving using wormhole-william
func NewClient(app fyne.App) *Client {
	return &Client{app: app}
}