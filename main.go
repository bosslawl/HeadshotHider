package main

import (
	"os"
	"path/filepath"
	"strings"
	"syscall"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/google/uuid"
	"golang.org/x/sys/windows"

	"HeadshotHider/core/assets"
	"HeadshotHider/core/gui"
	"HeadshotHider/core/util"
)

func main() {
	util.Logger.Println("Starting HeadshotHider...")
	currentName, err := os.Executable()
	if err != nil {
		util.Logger.Println("Failed to get executable name, error: " + err.Error())
		os.Exit(1)
	}
	newName := uuid.New().String()
	fileName := filepath.Base(currentName)

	if fileName == "HeadshotHider.exe" {
		err := os.Rename(currentName, newName+".exe")
		if err != nil {
			util.Logger.Println("Failed to rename executable, error: " + err.Error())
			os.Exit(1)
		}
		a := app.New()
		w := a.NewWindow("HeadshotHider")

		newMessage := widget.NewLabel("Please run the new executable: " + newName + ".exe")
		closeButton := widget.NewButton("Close", func() {
			a.Quit()
		})
		content := container.NewVBox(newMessage, closeButton)
		w.SetContent(content)
		w.Resize(fyne.NewSize(300, 100))
		w.SetMaster()
		w.CenterOnScreen()
		w.RequestFocus()
		w.ShowAndRun()
		return
	} else {
		if amAdmin() {
			a := app.NewWithID("bosslawl.HeadshotHider")
			assets.SetIcon(a)
			w := a.NewWindow("HeadshotHider")

			w.SetContent(gui.Create(a, w))
			w.Resize(fyne.NewSize(875, 500))
			w.SetMaster()
			w.CenterOnScreen()
			w.RequestFocus()
			w.ShowAndRun()
		} else {
			runAsAdmin()
		}
	}
}

func runAsAdmin() {
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
		util.Logger.Println("Failed to run as admin, error: " + err.Error())
		os.Exit(1)
	}
}

func amAdmin() bool {
	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	if err != nil {
		util.Logger.Println("Not running as admin, restarting as admin...")
		return false
	}
	return true
}
