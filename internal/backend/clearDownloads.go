package backend

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/bosslawl/HeadshotHider/v2/internal/util"
)

func ClearDownloads() error {
	downloadsDir := filepath.Join(os.Getenv("USERPROFILE"), "Downloads")
	files, _ := ioutil.ReadDir(downloadsDir)

	for _, file := range files {
		if strings.Contains(strings.ToLower(file.Name()), "headshot") || strings.Contains(strings.ToLower(file.Name()), "hs") {
			err := os.Remove(filepath.Join(downloadsDir, file.Name()))
			if err != nil {
				util.Logger.Println("Error deleting downloads:", err)
				return err
			}
		}
	}
	return nil
}
