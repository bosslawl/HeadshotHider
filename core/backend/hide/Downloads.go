package hide

import (
	"HeadshotHider/core/backend"
	"HeadshotHider/core/util"
	"os"
	"path/filepath"
	"strings"
)

func ClearDownloads() error {
	downloadsDir := backend.Directories["DownloadsDir"]
	files, err := os.ReadDir(downloadsDir)
	if err != nil {
		util.Logger.Printf("Error reading directory %s, error: %v\n", downloadsDir, err)
		return err
	}

	for _, file := range files {
		fileName := strings.ToLower(file.Name())
		if strings.Contains(fileName, "headshot") || strings.Contains(fileName, "hs") {
			err := os.Remove(filepath.Join(downloadsDir, file.Name()))
			if err != nil {
				util.Logger.Printf("Error deleting file %s, error: %v\n", file.Name(), err)
				return err
			}
		}
	}
	util.Logger.Println("Cleared downloads")
	return nil
}
