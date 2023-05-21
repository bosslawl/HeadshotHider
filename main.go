package main

import (
	//"fmt"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"

	//"time"

	"github.com/google/uuid"
	"golang.org/x/sys/windows"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

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

	if !amAdmin() && exe == "HeadshotHider.exe" {
		os.Rename(exe, uuid+".exe")
		cmd := exec.Command("cmd", "/C", "start", uuid+".exe")
		cmd.Run()
		runMeElevated()
		util.Logger.Println("Renaming executable and restarting as admin...")
		return
	} else if !amAdmin() && exe != "HeadshotHider.exe" {
		runMeElevated()
		util.Logger.Println("Restarting as admin...")
		return
	}

	a := app.NewWithID("io.github.bosslawl.HeadshotHider")
	assets.SetIcon(a)
	w := a.NewWindow("HeadshotHider")

	w.SetContent(ui.Create(a, w))
	w.Resize(fyne.NewSize(850, 500))
	w.SetMaster()
	util.Logger.Println("Running HeadshotHider")

	w.ShowAndRun()
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
