package backend

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/bosslawl/HeadshotHider/v2/internal/util"
)

/*
	Deletes all files in `Prefetch` in %windir%
	It may still log the name of this exe in the prefetch folder
	I have tried disabling prefetch via registry but it still logs
*/

func Prefetch() error {
	prefetch := "C:\\Windows\\Prefetch"
	files, _ := ioutil.ReadDir(prefetch)

	for _, file := range files {
		if strings.Contains(strings.ToLower(file.Name()), "headshot") || strings.Contains(strings.ToLower(file.Name()), "hs") {
			err := os.Remove(filepath.Join(prefetch, file.Name()))
			if err != nil {
				util.Logger.Println("Error deleting prefetch:", err)
				return err
			}
		}
	}
	return nil
}
