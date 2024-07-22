package hide

import (
	"HeadshotHider/core/backend"
	"HeadshotHider/core/util"
	"os"
	"path/filepath"
	"strings"
)

func ClearRecentFiles() error {
	files, err := os.ReadDir(backend.Directories["RecentDir"])
	if err != nil {
		util.Logger.Println("Error reading directory, error:", err)
		return err
	}

	for _, file := range files {
		fileName := strings.ToLower(file.Name())
		if strings.Contains(fileName, "headshot") || strings.Contains(fileName, "hs") {
			err := os.Remove(filepath.Join(backend.Directories["RecentDir"], file.Name()))
			if err != nil {
				util.Logger.Printf("Error deleting recent file %s, error: %v", file.Name(), err)
				return err
			}
		}
	}
	util.Logger.Println("Cleared recent files")
	return nil
}
