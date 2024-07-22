package hide

import (
	"HeadshotHider/core/backend"
	"HeadshotHider/core/util"
	"os"
	"path/filepath"
	"strings"
)

func ClearPrefetch() error {
	files, err := os.ReadDir(backend.Directories["PrefetchDir"])
	if err != nil {
		util.Logger.Println("Error reading directory, error:", err)
		return err
	}

	for _, file := range files {
		fileName := strings.ToLower(file.Name())
		if strings.Contains(fileName, "headshot") || strings.Contains(fileName, "hs") {
			err := os.Remove(filepath.Join(backend.Directories["PrefetchDir"], file.Name()))
			if err != nil {
				util.Logger.Printf("Error deleting prefetch file %s, error: %v", file.Name(), err)
				return err
			}
		}
	}
	util.Logger.Println("Cleared prefetch")
	return nil
}
