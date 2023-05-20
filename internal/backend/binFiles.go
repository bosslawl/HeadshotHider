package backend

import (
	"os"
)

/* 
	Deletes all files in `bin_files` in %temp% 
	This is just HSUWPHelper
*/

func BinFiles() {
	homeDir := UserHomeDir()

	bin := homeDir + "\\AppData\\Local\\Temp\\bin_files"
	os.RemoveAll(bin)
	os.Mkdir(bin, 0755)
}