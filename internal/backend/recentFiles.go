package backend

import (
	"os"
)

/*
	Deletes all files in `Recent` in %appdata%
*/

func RecentFiles() {
	recent := UserHomeDir() + "\\AppData\\Roaming\\Microsoft\\Windows\\Recent"
	os.RemoveAll(recent)
	os.Mkdir(recent, 0755)
}
