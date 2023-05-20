// Package transport handles sending and receiving using wormhole-william
package backend

import (
	"os"
	"runtime"

	"fyne.io/fyne/v2"
)

// Client defines the client for handling sending and receiving using wormhole-william
type Client struct {
	app fyne.App

	// DownloadPath holds the download path used for saving received files.
	DownloadPath string

	// LoaderPath holds the path to the loader used for sending files.
	LoaderPath string
}

// NewClient returns a new client for sending and receiving using wormhole-william
func NewClient(app fyne.App) *Client {
	return &Client{app: app}
}

func UserHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}
