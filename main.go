package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"HeadshotRemover/core"

	"github.com/google/uuid"
)

func main() {
	// if not elevated
	if !amAdmin() {
		fmt.Println("Please run as administrator... Please wait 5 seconds to relaunch")
		time.Sleep(5 * time.Second)
	}

	if _, err := os.Stat("HSLoader.exe"); os.IsNotExist(err) {
		fmt.Println("Please put this file in the same directory as HSLoader.exe")
	}

	name, err := os.Executable()
	if err != nil {
		panic(err)
	}
	uuid := uuid.New().String()
	exe := filepath.Base(name)
	os.Rename(exe, uuid+".exe")

	if exe == "HeadshotHider.exe" {
		cmd := exec.Command("cmd", "/C", "start", uuid+".exe")
		cmd.Run()
		os.Exit(0)
	}

	// if elevated, run the program
	core.CopyConfig()
	core.DeleteConfig()
	core.CopyRegistry()
	core.DeleteRegistry()
	core.Downloads()
	core.RecentFiles()
	core.BinFiles()
	core.BrowserHistory()
	core.RecyclingBin()
	core.LoaderUpdated()
	core.Prefetch()
}

func amAdmin() bool {
	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	if err != nil {
		return false
	}
	return true
}
