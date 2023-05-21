package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/bosslawl/HeadshotHider/v2/internal/backend"
)

type run struct {
	content *widget.Button

	client *backend.Client
	window fyne.Window
	canvas fyne.Canvas
	app    fyne.App
}

func newManage(a fyne.App, w fyne.Window, c *backend.Client) *run {
	return &run{app: a, window: w, client: c, canvas: w.Canvas()}
}

func update(progress *widget.ProgressBar) {
	progress.Value++
	progress.Refresh()
}

func (r *run) onRun() {
	progress := widget.NewProgressBar()
	progress.Min = 0
	progress.Max = 13
	progress.Value = 0

	container := container.NewVBox(progress)
	d := dialog.NewCustom("Clearing Data", "Please Wait...", container, r.window)
	d.Resize(fyne.NewSize(300, 100))
	d.Show()

	if r.client.Config() != nil {
		dialog.ShowError(r.client.Config(), r.window)
		return
	}
	update(progress)
	if r.client.DeleteLoader() != nil {
		dialog.ShowError(r.client.DeleteLoader(), r.window)
		return
	}
	update(progress)
	if backend.BinFiles() != nil {
		dialog.ShowError(backend.BinFiles(), r.window)
		return
	}
	update(progress)
	if backend.RecentFiles() != nil {
		dialog.ShowError(backend.RecentFiles(), r.window)
		return
	}
	update(progress)
	if backend.RecyclingBin() != nil {
		dialog.ShowError(backend.RecyclingBin(), r.window)
		return
	}
	update(progress)
	if r.client.RegistryKey() != nil {
		dialog.ShowError(r.client.RegistryKey(), r.window)
		return
	}
	update(progress)
	if backend.DeleteRegistryHistory() != nil {
		dialog.ShowError(backend.DeleteRegistryHistory(), r.window)
		return
	}
	update(progress)
	if backend.Downloads() != nil {
		dialog.ShowError(backend.Downloads(), r.window)
		return
	}
	update(progress)
	if backend.DeleteRegistryRun() != nil {
		dialog.ShowError(backend.DeleteRegistryRun(), r.window)
		return
	}
	update(progress)
	if backend.ARK() != nil {
		dialog.ShowError(backend.ARK(), r.window)
		return
	}
	update(progress)
	if backend.DeleteHistory() != nil {
		dialog.ShowError(backend.DeleteHistory(), r.window)
		return
	}
	update(progress)
	if backend.BrowserHistory() != nil {
		dialog.ShowError(backend.BrowserHistory(), r.window)
		return
	}
	update(progress)
	if backend.Prefetch() != nil {
		dialog.ShowError(backend.Prefetch(), r.window)
		return
	}
	update(progress)

	d.Hide()
	dialog.ShowInformation("HeadshotHider", "Successfully cleared all data.", r.window)
}

func (r *run) buildUI() *container.Scroll {

	r.content = &widget.Button{Text: "Execute", Icon: theme.DeleteIcon(), OnTapped: r.onRun}

	box := container.NewGridWithColumns(1, r.content)

	featuresBox := container.NewGridWithColumns(2,
		newBoldLabel("Copy & Delete Config (Set download path in settings)"),
		newBoldLabel("Delete Loader (Set loader path in settings)"),
		newBoldLabel("Copy & Delete Registry Key (Set download path in settings)"),
		newBoldLabel("Delete Registry History"),
		newBoldLabel("Delete Run History"),
		newBoldLabel("Empty Recycle Bin"),
		newBoldLabel("Delete Recent Files Opened"),
		newBoldLabel("Delete Any Files Pervading to HS In Downloads"),
		newBoldLabel("Close ARK"),
		newBoldLabel("Clear Prefetch"),
		newBoldLabel("Clear Browser History & Downloads (50/50 Chance)"),
		newBoldLabel("Clear Windows Defender History"),
	)

	return container.NewScroll(container.NewVBox(
		&widget.Card{Title: "Run", Content: box},
		&widget.Card{Title: "Features", Content: featuresBox},
	))
}

func (r *run) tabItem() *container.TabItem {
	return &container.TabItem{Text: "Manage", Icon: theme.FileIcon(), Content: r.buildUI()}
}
