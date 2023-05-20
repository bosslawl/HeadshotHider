package backend

import (
	"os/exec"
)

func ARK() {
	exec.Command("taskkill", "/IM", "shootergame.exe", "/F").Run()
}
