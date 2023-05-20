package ui

import (
	"errors"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	appearance "fyne.io/fyne/v2/cmd/fyne_settings/settings"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/bosslawl/HeadshotHider/v2/internal/backend"
	"github.com/bosslawl/HeadshotHider/v2/internal/util"
)

type settings struct {
	downloadPathEntry *widget.Entry
	loaderPathEntry   *widget.Entry

	client      *backend.Client
	preferences fyne.Preferences
	window      fyne.Window
	app         fyne.App
}

func newSettings(a fyne.App, w fyne.Window, c *backend.Client) *settings {
	return &settings{app: a, window: w, client: c, preferences: a.Preferences()}
}

func (s *settings) onDownloadsPathSubmitted(path string) {
	path = filepath.Clean(path)
	info, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		dialog.ShowInformation("Does not exist", "Please select a valid directory.", s.window)
		return
	} else if err != nil {
		fyne.LogError("Error when trying to verify directory", err)
		dialog.ShowError(err, s.window)
		return
	} else if !info.IsDir() {
		dialog.ShowInformation("Not a directory", "Please select a valid directory.", s.window)
		return
	}

	s.client.DownloadPath = path
	s.preferences.SetString("DownloadPath", s.client.DownloadPath)
	s.downloadPathEntry.SetText(s.client.DownloadPath)
}

func (s *settings) onLoaderPathSubmitted(path string) {
	path = filepath.Clean(path)
	info, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		dialog.ShowInformation("Does not exist", "Please select a valid directory.", s.window)
		return
	} else if err != nil {
		fyne.LogError("Error when trying to verify directory", err)
		dialog.ShowError(err, s.window)
		return
	} else if !info.IsDir() {
		dialog.ShowInformation("Not a directory", "Please select a valid directory.", s.window)
		return
	}

	s.client.LoaderPath = path
	s.preferences.SetString("LoaderPath", s.client.LoaderPath)
	s.loaderPathEntry.SetText(s.client.LoaderPath)
}

func (s *settings) onDownloadsPathSelected() {
	folder := dialog.NewFolderOpen(func(folder fyne.ListableURI, err error) {
		if err != nil {
			fyne.LogError("Error on selecting folder", err)
			dialog.ShowError(err, s.window)
			return
		} else if folder == nil {
			return
		}

		s.client.DownloadPath = folder.Path()
		s.preferences.SetString("DownloadPath", s.client.DownloadPath)
		s.downloadPathEntry.SetText(s.client.DownloadPath)
	}, s.window)

	folder.Resize(util.WindowSizeToDialog(s.window.Canvas().Size()))
	folder.Show()
}

func (s *settings) onLoaderPathSelected() {
	folder := dialog.NewFolderOpen(func(folder fyne.ListableURI, err error) {
		if err != nil {
			fyne.LogError("Error on selecting folder", err)
			dialog.ShowError(err, s.window)
			return
		} else if folder == nil {
			return
		}

		s.client.LoaderPath = folder.Path()
		s.preferences.SetString("LoaderPath", s.client.LoaderPath)
		s.loaderPathEntry.SetText(s.client.LoaderPath)
	}, s.window)

	folder.Resize(util.WindowSizeToDialog(s.window.Canvas().Size()))
	folder.Show()
}

// getPreferences is used to set the preferences on startup without saving at the same time.
func (s *settings) getPreferences() {
	s.client.DownloadPath = s.preferences.StringWithFallback("DownloadPath", util.UserDownloadsFolder())
	s.client.LoaderPath = s.preferences.StringWithFallback("LoaderPath", util.UserDownloadsFolder())
	s.downloadPathEntry.Text = s.client.DownloadPath
	s.loaderPathEntry.Text = s.client.LoaderPath
}

func (s *settings) buildUI() *container.Scroll {

	pathSelector := &widget.Button{Icon: theme.FolderOpenIcon(), Importance: widget.LowImportance, OnTapped: s.onDownloadsPathSelected}
	s.downloadPathEntry = &widget.Entry{Wrapping: fyne.TextTruncate, OnSubmitted: s.onDownloadsPathSubmitted, ActionItem: pathSelector}

	loaderSelector := &widget.Button{Icon: theme.FolderOpenIcon(), Importance: widget.LowImportance, OnTapped: s.onLoaderPathSelected}
	s.loaderPathEntry = &widget.Entry{Wrapping: fyne.TextTruncate, OnSubmitted: s.onLoaderPathSubmitted, ActionItem: loaderSelector}

	s.getPreferences()

	interfaceContainer := appearance.NewSettings().LoadAppearanceScreen(s.window)

	dataContainer := container.NewGridWithColumns(2,
		newBoldLabel("Download files to"), s.downloadPathEntry,
		newBoldLabel("Loader path"), s.loaderPathEntry,
	)

	return container.NewScroll(container.NewVBox(
		&widget.Card{Title: "User Interface", Content: interfaceContainer},
		&widget.Card{Title: "Data Handling", Content: dataContainer},
	))
}

func (s *settings) tabItem() *container.TabItem {
	return &container.TabItem{Text: "Settings", Icon: theme.SettingsIcon(), Content: s.buildUI()}
}

func newBoldLabel(text string) *widget.Label {
	return &widget.Label{Text: text, TextStyle: fyne.TextStyle{Bold: true}}
}
