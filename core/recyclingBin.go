package core

import (
	"log"
	"os/exec"
)

/*
	Empties the recycling bin 
	Uses powershell to run the command
*/

func RecyclingBin() {
	cmd := exec.Command("PowerShell", "-Command", "Clear-RecycleBin -Force")
	cmd.Run()
	log.Println("Recycling bin emptied successfully")
}
