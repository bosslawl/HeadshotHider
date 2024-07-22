package hide

import (
	"HeadshotHider/core/util"

	"golang.org/x/sys/windows/registry"
)

func ClearUserAssist() error {
	key, _ := registry.OpenKey(registry.CURRENT_USER, `SOFTWARE\Software\Microsoft\Windows\CurrentVersion\Explorer\UserAssist`, registry.ALL_ACCESS)
	err := DeleteSingle(key)
	if err != nil {
		util.Logger.Println("Error clearing user assist:", err)
		return err
	}
	return nil
}
