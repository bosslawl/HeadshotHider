package hide

import (
	"HeadshotHider/core/backend"
	"HeadshotHider/core/util"
	"os"
	"path/filepath"
	"strings"
)

func ClearCrashDumps() error {
	crashDir := backend.Directories["CrashDir"]
	files, err := os.ReadDir(crashDir)
	if err != nil {
		util.Logger.Printf("Error reading directory %s, error: %v\n", crashDir, err)
		return err
	}

	for _, file := range files {
		fileName := strings.ToLower(file.Name())
		if strings.Contains(fileName, "headshot") || strings.Contains(fileName, "hs") {
			err := os.Remove(filepath.Join(crashDir, file.Name()))
			if err != nil {
				util.Logger.Printf("Error deleting file %s, error: %v\n", file.Name(), err)
				return err
			}
		}
	}
	util.Logger.Println("Cleared crash dumps")
	return nil
}
