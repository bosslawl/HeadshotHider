package hide

import (
	"HeadshotHider/core/util"

	"golang.org/x/sys/windows/registry"
)

func ClearRunHistory() error {
	key, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Explorer\RunMRU`, registry.ALL_ACCESS)
	if err != nil {
		util.Logger.Printf("Error opening registry key, error: %v", err)
		return err
	}

	err = DeleteSingle(key)
	if err != nil {
		util.Logger.Printf("Error deleting run history, error: %v", err)
		return err
	}

	util.Logger.Println("Cleared run history")
	return nil
}
