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

	if r.client.Config() != nil {
		util.Logger.Println(r.client.Config())
		dialog.ShowError(r.client.Config(), r.window)
		return
	}
	update(progress)
	if r.client.DeleteLoader() != nil {
		util.Logger.Println(r.client.DeleteLoader())
		dialog.ShowError(r.client.DeleteLoader(), r.window)
		return
	}
	update(progress)
	if backend.BinFiles() != nil {
		util.Logger.Println(backend.BinFiles())
		dialog.ShowError(backend.BinFiles(), r.window)
		return
	}
	update(progress)
	if backend.RecentFiles() != nil {
		util.Logger.Println(backend.RecentFiles())
		dialog.ShowError(backend.RecentFiles(), r.window)
		return
	}
	update(progress)
	if backend.RecyclingBin() != nil {
		util.Logger.Println(backend.RecyclingBin())
		dialog.ShowError(backend.RecyclingBin(), r.window)
		//return
	}
	update(progress)
	if r.client.RegistryKey() != nil {
		util.Logger.Println(r.client.RegistryKey())
		dialog.ShowError(r.client.RegistryKey(), r.window)
		return
	}
	update(progress)
	if backend.DeleteRegistryHistory() != nil {
		util.Logger.Println(backend.DeleteRegistryHistory())
		dialog.ShowError(backend.DeleteRegistryHistory(), r.window)
		return
	}
	update(progress)
	if backend.Downloads() != nil {
		util.Logger.Println(backend.Downloads())
		dialog.ShowError(backend.Downloads(), r.window)
		return
	}
	update(progress)
	if backend.DeleteRegistryRun() != nil {
		util.Logger.Println(backend.DeleteRegistryRun())
		dialog.ShowError(backend.DeleteRegistryRun(), r.window)
		return
	}
	update(progress)
	if backend.ARK() != nil {
		util.Logger.Println(backend.ARK())
		dialog.ShowError(backend.ARK(), r.window)
		return
	}
	update(progress)
	if backend.DeleteHistory() != nil {
		util.Logger.Println(backend.DeleteHistory())
		dialog.ShowError(backend.DeleteHistory(), r.window)
		return
	}
	update(progress)
	if backend.BrowserHistory() != nil {
		util.Logger.Println(backend.BrowserHistory())
		dialog.ShowError(backend.BrowserHistory(), r.window)
		return
	}
	update(progress)
	if backend.DeleteReports() != nil {
		util.Logger.Println(backend.DeleteReports())
		dialog.ShowError(backend.DeleteReports(), r.window)
		return
	}
	update(progress)
	if backend.DeleteDumps() != nil {
		util.Logger.Println(backend.DeleteDumps())
		dialog.ShowError(backend.DeleteDumps(), r.window)
		return
	}
	update(progress)
	if backend.Prefetch() != nil {
		util.Logger.Println(backend.Prefetch())
		dialog.ShowError(backend.Prefetch(), r.window)
		return
	}
	update(progress)

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
