package main

import (
	//"fmt"

	"fmt"
	"os"
	"path/filepath"
	"strings"
	"syscall"

	//"time"

	"github.com/google/uuid"
	"golang.org/x/sys/windows"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/bosslawl/HeadshotHider/v2/internal/assets"
	"github.com/bosslawl/HeadshotHider/v2/internal/ui"
	"github.com/bosslawl/HeadshotHider/v2/internal/util"
)

func main() {
	util.Logger.Println("Starting HeadshotHider...")
	name, err := os.Executable()
	if err != nil {
		panic(err)
	}
	uuid := uuid.New().String()
	exe := filepath.Base(name)

	if exe == "HeadshotHider.exe" {
		err := os.Rename(exe, uuid+".exe")
		if err != nil {
			util.Logger.Println("Error renaming executable:", err)
			return
		}
		util.Logger.Println("Renamed executable")

		a := app.New()
		w := a.NewWindow("HeadshotHider")
		errorLabel := widget.NewLabel("Please run the new executable: " + uuid + ".exe")
		closeButton := widget.NewButton("Close", func() {
			a.Quit()
		})
		content := container.NewVBox(errorLabel, closeButton)
		w.SetContent(content)
		w.Resize(fyne.NewSize(300, 100))
		w.ShowAndRun()
		return
	} else {
		if amAdmin() {
			a := app.NewWithID("io.github.bosslawl.HeadshotHider")
			assets.SetIcon(a)
			w := a.NewWindow("HeadshotHider")

			w.SetContent(ui.Create(a, w))
			w.Resize(fyne.NewSize(875, 500))
			w.SetMaster()
			util.Logger.Println("Running HeadshotHider")

			w.ShowAndRun()
		} else {
			runMeElevated()
		}
	}
}

func runMeElevated() {
	verb := "runas"
	exe, _ := os.Executable()
	cwd, _ := os.Getwd()
	args := strings.Join(os.Args[1:], " ")

	verbPtr, _ := syscall.UTF16PtrFromString(verb)
	exePtr, _ := syscall.UTF16PtrFromString(exe)
	cwdPtr, _ := syscall.UTF16PtrFromString(cwd)
	argPtr, _ := syscall.UTF16PtrFromString(args)

	var showCmd int32 = 1 //SW_NORMAL

	err := windows.ShellExecute(0, verbPtr, exePtr, argPtr, cwdPtr, showCmd)
	if err != nil {
		fmt.Println(err)
	}
}

func amAdmin() bool {
	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	if err != nil {
		return false
	}
	return true
}
