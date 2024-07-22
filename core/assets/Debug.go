package assets

import (
	_ "embed"

	"fyne.io/fyne/v2"
)

//go:embed icon/icon-512.png
var iconData []byte

func SetIcon(a fyne.App) {
	a.SetIcon(
		&fyne.StaticResource{
			StaticName:    "icon.png",
			StaticContent: iconData,
		},
	)
}
