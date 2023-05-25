package backend

import (
	"os"
	"runtime"

	"fyne.io/fyne/v2"
)

type Client struct {
	app fyne.App

	// DownloadPath holds the download path used for saving received files.
	DownloadPath string

	// LoaderPath holds the path to the loader used for sending files.
	LoaderPath string
}

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
