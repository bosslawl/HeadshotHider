package backend

import (
	"os"

	"github.com/bosslawl/HeadshotHider/v2/internal/util"
)

/*
	Deletes all files in `bin_files` in %temp%
	This is just HSUWPHelper
*/

func ClearBinFiles() error {
	homeDir := UserHomeDir()

	bin := homeDir + "\\AppData\\Local\\Temp\\bin_files"
	err := os.RemoveAll(bin)
	if err != nil {
		util.Logger.Println("Error deleting bin files:", err)
		return err
	}
	return nil
}
