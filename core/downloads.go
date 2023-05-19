package core

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

/*
	Deletes all files in `Downloads` that contain `headshot` or `hs`
*/

func Downloads() {
	downloadsDir := filepath.Join(os.Getenv("USERPROFILE"), "Downloads")
	files, err := ioutil.ReadDir(downloadsDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if strings.Contains(strings.ToLower(file.Name()), "headshot") || strings.Contains(strings.ToLower(file.Name()), "hs") {
			err := os.Remove(filepath.Join(downloadsDir, file.Name()))
			if err != nil {
				log.Fatal(err)
			}
			log.Println("File deleted:", file.Name())
		}
	}
}
