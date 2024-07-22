package gui

import (
	"HeadshotHider/core/backend/hide"
	"HeadshotHider/core/util"
	"errors"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	appearance "fyne.io/fyne/v2/cmd/fyne_settings/settings"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type settings struct {
	downloadFilePath *widget.Entry
	loaderFilePath   *widget.Entry

	client      *hide.Client
	window      fyne.Window
	app         fyne.App
	preferences fyne.Preferences
}

func Settings(a fyne.App, w fyne.Window, c *hide.Client) *settings {
	return &settings{app: a, window: w, client: c, preferences: a.Preferences()}
}

func (s *settings) DownloadsPath(path string) {
	path = filepath.Clean(path)
	check, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		dialog.ShowInformation("The path does not exist.", "Please select a valid directory.", s.window)
		return
	} else if !check.IsDir() {
		dialog.ShowInformation("The path is not a directory.", "Please select a valid directory.", s.window)
		return
	} else if err != nil {
		dialog.ShowError(err, s.window)
		util.Logger.Println("Error setting downloads path:", err.Error())
		return
	}

	s.client.DownloadPath = path
	s.preferences.SetString("DownloadPath", path)
	s.downloadFilePath.SetText(s.client.DownloadPath)
}

func (s *settings) LoaderPath(path string) {
	path = filepath.Clean(path)
	check, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		dialog.ShowInformation("The path does not exist.", "Please select a valid directory.", s.window)
		return
	} else if !check.IsDir() {
		dialog.ShowInformation("The path is not a directory.", "Please select a valid directory.", s.window)
		return
	} else if err != nil {
		dialog.ShowError(err, s.window)
		util.Logger.Println("Error setting loader path:", err.Error())
		return
	}

	s.client.LoaderPath = path
	s.preferences.SetString("LoaderPath", path)
	s.loaderFilePath.SetText(s.client.LoaderPath)
}

func (s *settings) DownloadsPathSelect() {
	folder := dialog.NewFolderOpen(func(folder fyne.ListableURI, err error) {
		if err != nil {
			dialog.ShowError(err, s.window)
			util.Logger.Println("Error selecting downloads path:", err.Error())
			return
		} else if folder == nil {
			return
		}

		s.client.DownloadPath = folder.Path()
		s.preferences.SetString("DownloadPath", s.client.DownloadPath)
		s.downloadFilePath.SetText(s.client.DownloadPath)
	}, s.window)

	folder.Resize(util.WindowSizeToDialog(s.window.Canvas().Size()))
	folder.Show()
}

func (s *settings) LoaderPathSelect() {
	folder := dialog.NewFolderOpen(func(folder fyne.ListableURI, err error) {
		if err != nil {
			dialog.ShowError(err, s.window)
			util.Logger.Println("Error selecting downloads path:", err.Error())
			return
		} else if folder == nil {
			return
		}

		s.client.LoaderPath = folder.Path()
		s.preferences.SetString("LoaderPath", s.client.LoaderPath)
		s.loaderFilePath.SetText(s.client.LoaderPath)
	}, s.window)

	folder.Resize(util.WindowSizeToDialog(s.window.Canvas().Size()))
	folder.Show()
}

func (s *settings) GetPreferences() {
	s.client.DownloadPath = s.preferences.StringWithFallback("DownloadPath", util.UserDownloadsFolder())
	s.client.LoaderPath = s.preferences.StringWithFallback("LoaderPath", util.UserDownloadsFolder())
	s.downloadFilePath.Text = s.client.DownloadPath
	s.loaderFilePath.Text = s.client.LoaderPath
}

func newBoldLabel(text string) *widget.Label {
	return &widget.Label{Text: text, TextStyle: fyne.TextStyle{Bold: true}}
}

func (s *settings) BuildUI() *container.Scroll {
	downloadsSelect := &widget.Button{Icon: theme.FolderOpenIcon(), Importance: widget.LowImportance, OnTapped: s.DownloadsPathSelect}
	s.downloadFilePath = &widget.Entry{Wrapping: fyne.TextTruncate, OnSubmitted: s.DownloadsPath, ActionItem: downloadsSelect}

	loaderSelect := &widget.Button{Icon: theme.FolderOpenIcon(), Importance: widget.LowImportance, OnTapped: s.LoaderPathSelect}
	s.loaderFilePath = &widget.Entry{Wrapping: fyne.TextTruncate, OnSubmitted: s.LoaderPath, ActionItem: loaderSelect}

	s.GetPreferences()

	gui := appearance.NewSettings().LoadAppearanceScreen(s.window)

	dataContainer := container.NewGridWithRows(2,
		newBoldLabel("Download Path"), s.downloadFilePath,
		newBoldLabel("Loader Path"), s.loaderFilePath,
	)

	return container.NewScroll(container.NewVBox(
		&widget.Card{Title: "GUI", Content: gui},
		&widget.Card{Title: "Data Handling", Content: dataContainer},
	))
}

func (s *settings) TabItem() *container.TabItem {
	return &container.TabItem{Text: "Settings", Icon: theme.SettingsIcon(), Content: s.BuildUI()}
}
