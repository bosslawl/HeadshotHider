package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/bosslawl/HeadshotHider/v2/internal/backend"
	"github.com/bosslawl/HeadshotHider/v2/internal/util"
)

type manage func() error

type manageItem struct {
	name      string
	component manage
}

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
	progress.Max = 15
	progress.Value = 0

	container := container.NewVBox(progress)
	d := dialog.NewCustom("Clearing Data", "Please Wait...", container, r.window)
	d.Resize(fyne.NewSize(300, 100))
	d.Show()

	var items = []manageItem{
		{"Clearing Config", r.client.ClearConfig},
		{"Clearing Delete Loader", r.client.DeleteLoader},
		{"Clearing Bin Files", backend.ClearBinFiles},
		{"Clearing Recent Files", backend.ClearRecentFiles},
		{"Clearing Recycling Bin", backend.ClearRecyclingbin},
		{"Clearing Registry Key", r.client.ClearRegistryKey},
		{"Clearing Registry History", backend.ClearStoreHistory},
		{"Clearing Registry Run", backend.ClearRunHistory},
		{"Clearing Downloads", backend.ClearDownloads},
		{"Closing ARK", backend.CloseARK},
		{"Clearing Windows Defender", backend.ClearWindowsDefender},
		{"Clearing Browser History", backend.ClearBrowserHistory},
		{"Clearing Crash Dumps", backend.ClearCrashDumps},
		{"Clearing Reports", backend.ClearCrashReports},
		{"Clearing Prefetch", backend.ClearPrefetch},
	}

	for _, item := range items {
		err := item.component()
		if err != nil {
			util.Logger.Println("Error clearing data:", item.name)
			dialog.ShowError(err, r.window)
		}
		update(progress)
	}

	d.Hide()
	util.Logger.Println("Successfully cleared all data.")
	dialog.ShowInformation("HeadshotHider", "Successfully cleared all data.", r.window)
}

func (r *run) buildUI() *container.Scroll {

	r.content = &widget.Button{Text: "Execute", Icon: theme.DeleteIcon(), OnTapped: r.onRun}

	box := container.NewGridWithColumns(1, r.content)

	featuresBox := container.NewGridWithColumns(2,
		newBoldLabel("- Copy & Delete Config (Set download path in settings)"),
		newBoldLabel("- Delete Loader (Set loader path in settings)"),
		newBoldLabel("- Copy & Delete Registry Key (Set download path in settings)"),
		newBoldLabel("- Delete Registry History"),
		newBoldLabel("- Delete Run History"),
		newBoldLabel("- Empty Recycle Bin"),
		newBoldLabel("- Delete Recent Files Opened"),
		newBoldLabel("- Delete Any Files Pervading to HS In Downloads"),
		newBoldLabel("- Close ARK"),
		newBoldLabel("- Clear Prefetch"),
		newBoldLabel("- Clear Browser History & Downloads (50/50 Chance)"),
		newBoldLabel("- Clear Windows Defender History"),
		newBoldLabel("- Clear Windows Crash Dumps"),
		newBoldLabel("- Clear Windows Dump Reports"),
	)

	miscBox := container.NewGridWithColumns(2,
		newBoldLabel("- You will still need to delete emails pervading to HS"),
		newBoldLabel("- You will still need to delete any sellix orders pervading to HS"),
	)

	return container.NewScroll(container.NewVBox(
		&widget.Card{Title: "Run", Content: box},
		&widget.Card{Title: "Features", Content: featuresBox},
		&widget.Card{Title: "Misc", Content: miscBox},
	))
}

func (r *run) tabItem() *container.TabItem {
	return &container.TabItem{Text: "Manage", Icon: theme.FileIcon(), Content: r.buildUI()}
}
