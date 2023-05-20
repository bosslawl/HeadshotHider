package backend

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func Downloads() {
	downloadsDir := filepath.Join(os.Getenv("USERPROFILE"), "Downloads")
	files, _ := ioutil.ReadDir(downloadsDir)

	for _, file := range files {
		if strings.Contains(strings.ToLower(file.Name()), "headshot") || strings.Contains(strings.ToLower(file.Name()), "hs") {
			os.Remove(filepath.Join(downloadsDir, file.Name()))
		}
	}
}
