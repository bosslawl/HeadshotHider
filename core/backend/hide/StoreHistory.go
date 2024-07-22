package hide

import (
	"HeadshotHider/core/util"

	"golang.org/x/sys/windows/registry"
)

func ClearStoreHistory() error {
	key, _ := registry.OpenKey(registry.CURRENT_USER, `SOFTWARE\Microsoft\Windows NT\CurrentVersion\AppCompatFlags\Compatibility Assistant\Store`, registry.ALL_ACCESS)
	err := DeleteSingle(key)
	if err != nil {
		util.Logger.Println("Error deleting registry history:", err)
		return err
	}
	return nil
}
