package backend

import (
	"os/exec"

	"github.com/bosslawl/HeadshotHider/v2/internal/util"
)

/*
	Empties the recycling bin
	Uses powershell to run the command
*/

func RecyclingBin() error {
	cmd := exec.Command("PowerShell", "-Command", "Clear-RecycleBin -Force")
	err := cmd.Run()
	if err != nil {
		util.Logger.Println("Error emptying recycling bin:", err)
		//return err
	}
	return nil
}
