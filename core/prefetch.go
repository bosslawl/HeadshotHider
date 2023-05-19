package core

import (
	"log"
	"os"
)

/*
	Deletes all files in `Prefetch` in %windir%
	It may still log the name of this exe in the prefetch folder
	I have tried disabling prefetch via registry but it still logs 
*/

func Prefetch() {
	prefetch := "C:\\Windows\\Prefetch"

	os.RemoveAll(prefetch)
	log.Println("Prefetch files removed from:", prefetch)
}
