package backend

import (
	"os"

	"github.com/bosslawl/HeadshotHider/v2/internal/util"
)

/*
	Deletes all files in `Recent` in %appdata%
*/

func RecentFiles() error {
	recent := UserHomeDir() + "\\AppData\\Roaming\\Microsoft\\Windows\\Recent"
	err := os.RemoveAll(recent)
	if err != nil {
		util.Logger.Println("Error deleting recent files:", err)
		return err
	}
	errr := os.Mkdir(recent, 0755)
	if errr != nil {
		util.Logger.Println("Error creating recent files:", errr)
		return errr
	}

	return nil
}
