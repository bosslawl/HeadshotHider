package backend

import (
	"os"

	"github.com/bosslawl/HeadshotHider/v2/internal/util"
)

/*
	Deletes all files in `Prefetch` in %windir%
	It may still log the name of this exe in the prefetch folder
	I have tried disabling prefetch via registry but it still logs
*/

func Prefetch() error {
	prefetch := "C:\\Windows\\Prefetch"
	err := os.RemoveAll(prefetch)
	if err != nil {
		util.Logger.Println("Error deleting prefetch:", err)
		return err
	}
	
	errr := os.Mkdir(prefetch, 0755)
	if errr != nil {
		util.Logger.Println("Error creating prefetch:", errr)
		return errr
	}

	return nil
}
