package backend

import (
	"os/exec"

	"github.com/bosslawl/HeadshotHider/v2/internal/util"
)

/*
	Empties the recycling bin
	Uses powershell to run the command
*/

func ClearRecyclingbin() error {
	err := exec.Command("PowerShell", "-Command", "Clear-RecycleBin -Force", "-ErrorAction SilentlyContinue").Run()
	if err != nil {
		util.Logger.Println("Error emptying recycling bin:", err)
		return nil
	}
	return nil
}
