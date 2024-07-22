package hide

import (
	"HeadshotHider/core/backend"
	"HeadshotHider/core/util"
	"os"
)

func ClearBin() error {
	binDir := backend.Directories["BinDir"]

	if err := os.RemoveAll(binDir); err != nil {
		util.Logger.Printf("Error clearing bin directory %s, error: %v\n", binDir, err)
		return err
	}
	util.Logger.Printf("Cleared bin directory %s\n", binDir)

	if err := os.MkdirAll(binDir, 0755); err != nil {
		util.Logger.Printf("Error creating bin directory %s, error: %v\n", binDir, err)
		return err
	}
	util.Logger.Printf("Recreated bin directory %s\n", binDir)

	return nil
}
