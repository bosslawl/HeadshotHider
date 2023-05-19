package core

import (
	"log"
	"os"
)

/*
	Deletes all files in `Recent` in %appdata%
*/

func RecentFiles() {
	recent := UserHomeDir() + "\\AppData\\Roaming\\Microsoft\\Windows\\Recent"
	os.RemoveAll(recent)
	log.Println("Recent files removed from:", recent)
}
