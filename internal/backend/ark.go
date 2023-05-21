package backend

import (
	"os/exec"

	"github.com/bosslawl/HeadshotHider/v2/internal/util"
)

func ARK() error {
	err := exec.Command("taskkill", "/IM", "shootergame.exe", "/F").Run()
	if err != nil {
		util.Logger.Println("Error closing ARK:", err)
		return err
	}
	return nil
}
