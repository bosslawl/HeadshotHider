package backend

import (
	"golang.org/x/sys/windows/registry"
)

func DeleteRegistryRun() {
	key, _ := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Explorer\RunMRU`, registry.ALL_ACCESS)
	DeleteSingle(key)
}