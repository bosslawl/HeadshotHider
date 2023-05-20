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

func (r *run) onRun() {
	backend.BinFiles()
	backend.RecentFiles()
	backend.RecyclingBin()
	r.client.RegistryKey()
	backend.DeleteRegistryHistory()
	backend.Downloads()
	backend.DeleteRegistryRun()
	backend.ARK()
	r.client.DeleteLoader()
	r.client.Config()
	backend.DeleteHistory()
	backend.BrowserHistory()
	backend.Prefetch()

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
