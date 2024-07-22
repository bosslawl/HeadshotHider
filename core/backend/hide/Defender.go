package hide

import (
	"HeadshotHider/core/backend"
	"HeadshotHider/core/util"
	"os"
)

func ClearWindowsDefender() error {
	defendDir := backend.Directories["DefendDir"]

	err := os.RemoveAll(defendDir)
	if err != nil {
		util.Logger.Printf("Error deleting defender history directory %s, error: %v\n", defendDir, err)
		return err
	}

	err = os.MkdirAll(defendDir, 0755)
	if err != nil {
		util.Logger.Printf("Error recreating defender history directory %s, error: %v\n", defendDir, err)
		return err
	}

	util.Logger.Printf("Cleared defender history in directory %s\n", defendDir)
	return nil
}
