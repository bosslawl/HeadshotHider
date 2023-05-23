package backend

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/bosslawl/HeadshotHider/v2/internal/util"
)

func DeleteDumps() error {
	files, _ := ioutil.ReadDir(UserHomeDir() + "\\AppData\\Local\\CrashDumps")

	for _, file := range files {
		if strings.Contains(strings.ToLower(file.Name()), "headshot") || strings.Contains(strings.ToLower(file.Name()), "hs") {
			err := os.Remove(filepath.Join(UserHomeDir() + "\\Appdata\\Local\\CrashDumps", file.Name()))
			if err != nil {
				util.Logger.Println("Error deleting crash dumps:", err)
				return err
			}
		}
	}
	return nil
}
