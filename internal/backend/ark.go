package backend

import (
	"os/exec"
)

func ARK() error {
	exec.Command("taskkill", "/IM", "shootergame.exe", "/F").Run()
	return nil
}
