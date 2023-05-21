package backend

import (
	"github.com/bosslawl/HeadshotHider/v2/internal/util"
	"golang.org/x/sys/windows/registry"
)

func DeleteRegistryRun() error {
	key, _ := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Explorer\RunMRU`, registry.ALL_ACCESS)
	err := DeleteSingle(key)
	if err != nil {
		util.Logger.Println("Error deleting registry run:", err)
		return err
	}
	return nil
}
