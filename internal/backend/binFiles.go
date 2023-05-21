package backend

import (
	"os"
)

/* 
	Deletes all files in `bin_files` in %temp% 
	This is just HSUWPHelper
*/

func BinFiles() error {
	homeDir := UserHomeDir()

	bin := homeDir + "\\AppData\\Local\\Temp\\bin_files"
	os.RemoveAll(bin)
	return nil
}