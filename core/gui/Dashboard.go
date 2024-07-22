package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"HeadshotHider/core/backend/hide"
	"HeadshotHider/core/util"
)

type Manage func() error

type ManageItem struct {
	name      string
	component Manage
}

type Main struct {
	content *widget.Button

	client   *hide.Client
	window   fyne.Window
	canvas   fyne.Canvas
	app      fyne.App
}

func Dashboard(a fyne.App, w fyne.Window, c *hide.Client) *Main {
	return &Main{app: a, window: w, client: c, canvas: w.Canvas()}
}

func Update(progress *widget.ProgressBar) {
	progress.Value++
	progress.Refresh()
}

func (m *Main) Hide() {
	progress := widget.NewProgressBar()
	progress.Min = 0
	progress.Max = 15
	progress.Value = 0

	container := container.NewVBox(progress)
	d := dialog.NewCustom("Clearing Data", "Please Wait...", container, m.window)
	d.Resize(fyne.NewSize(300, 100))
	d.Show()

	var items = []ManageItem{
		{"Clearing Config", m.client.ClearConfig},
		{"Clearing Loader", m.client.ClearLoader},
		{"Clearing SerialKey", m.client.ClearRegistryKey},
		{"Clearing Bin", hide.ClearBin},
		{"Clearing Registry", hide.ClearStoreHistory},
		{"Clearing Crash Dumps", hide.ClearCrashDumps},
		{"Clearing Reports", hide.ClearCrashReports},
		{"Clearing Windows Defender", hide.ClearWindowsDefender},
		{"Clearing Downloads", hide.ClearDownloads},
		{"Clearing Recent", hide.ClearRecentFiles},
		{"Clearing Run History", hide.ClearRunHistory},
		{"Clearing Browser History", hide.ClearHistory},
		{"Clearing Recycling Bin", hide.ClearRecyclingbin},
		{"Closing ARK", hide.CloseARK},
		{"Clearing Prefetch", hide.ClearPrefetch},
		{"Clearing UserAssist", hide.ClearUserAssist},
	}

	for _, item := range items {
		err := item.component()
		if err != nil {
			util.Logger.Println("Error clearing data, error:", item.name)
			dialog.ShowError(err, m.window)
		}
		Update(progress)
	}

	d.Hide()
	util.Logger.Println("Successfully cleared all data.")
	dialog.ShowInformation("HeadshotHider", "Successfully cleared all data.", m.window)
}

func (m *Main) Restore() {
	progress := widget.NewProgressBar()
	progress.Min = 0
	progress.Max = 3
	progress.Value = 0

	container := container.NewVBox(progress)
	d := dialog.NewCustom("Restoring Data", "Please Wait...", container, m.window)
	d.Resize(fyne.NewSize(300, 100))
	d.Show()

	var items = []ManageItem{

	}

	for _, item := range items {
		err := item.component()
		if err != nil {
			util.Logger.Println("Error restoring data, error:", item.name)
			dialog.ShowError(err, m.window)
		}
		Update(progress)
	}

	d.Hide()
	util.Logger.Println("Successfully restored all data.")
	dialog.ShowInformation("HeadshotHider", "Successfully restored all data.", m.window)
}

func (m *Main) BuildUI() *container.Scroll {
	m.content = &widget.Button{Text: "Execute", OnTapped: m.Hide, Icon: theme.FileIcon(), IconPlacement: widget.ButtonIconLeadingText}
	hideBox := container.NewGridWithColumns(1, m.content)

	m.content = &widget.Button{Text: "Execute", OnTapped: m.Restore, Icon: theme.FileIcon(), IconPlacement: widget.ButtonIconLeadingText}
	restoreBox := container.NewGridWithColumns(1, m.content)

	dataContainer := container.NewGridWithRows(2,
		newBoldLabel("Hide"), hideBox,
		newBoldLabel("Restore"), restoreBox,
	)

	return container.NewScroll(container.NewVBox(
		&widget.Card{Title: "Functions", Content: dataContainer},
	))
}

func (m *Main) TabItem() *container.TabItem {
	return &container.TabItem{Text: "Dashboard", Icon: theme.HomeIcon(), Content: m.BuildUI()}
}
