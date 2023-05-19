package core

import (
	"log"
	"os"
)

/* 
	Deletes all files in `bin_files` in %temp% 
	This is just HSUWPHelper
*/

func BinFiles() {
	homeDir := UserHomeDir()

	bin := homeDir + "\\AppData\\Local\\Temp\\bin_files"
	binerr := os.RemoveAll(bin)
	if binerr != nil {
		log.Fatalf("%v", binerr)
		return
	}
	log.Println("Bin files deleted from:", bin)
}
