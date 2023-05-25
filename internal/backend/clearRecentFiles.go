package backend

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/bosslawl/HeadshotHider/v2/internal/util"
)

/*
	Deletes all files in `Recent` in %appdata%
*/

func ClearRecentFiles() error {
	recent := UserHomeDir() + "\\AppData\\Roaming\\Microsoft\\Windows\\Recent"
	files, _ := ioutil.ReadDir(recent)

	for _, file := range files {
		if strings.Contains(strings.ToLower(file.Name()), "headshot") || strings.Contains(strings.ToLower(file.Name()), "hs") {
			err := os.Remove(filepath.Join(recent, file.Name()))
			if err != nil {
				util.Logger.Println("Error deleting recent files:", err)
				return err
			}
		}
	}
	return nil
}
