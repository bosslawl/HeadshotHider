package hide

import (
	"HeadshotHider/core/util"
	"os/exec"
)

func ClearRecyclingbin() error {
	cmd := exec.Command("PowerShell", "-Command", "Clear-RecycleBin", "-Force", "-ErrorAction", "SilentlyContinue")
	err := cmd.Run()
	if err != nil {
		util.Logger.Printf("Error emptying recycling bin, error: %v", err)
		return err
	}
	util.Logger.Println("Emptied recycling bin")
	return nil
}
