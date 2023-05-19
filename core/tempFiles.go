package core

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	//"path/filepath"

	"github.com/google/uuid"
)

/*
	Starts off by opening the config file and copying it
	Then asks the user where to dump the config and dumps it in a uuid file
	Then deletes the config file
*/

func CopyConfig() {
	homeDir := UserHomeDir()
	TempFiles := homeDir + "\\AppData\\Local\\Packages\\StudioWildcard.4558480580BB9_1w2mm55455e38\\AC\\Temp"
	//HSConfigs := TempFiles + "\\HeadshotConfigs"
	var newpath string

	fmt.Println(`Choose where to save your config to so you don't loose it.`)
	fmt.Println(`This is your full headshot config, so make sure you save it in a safe place.`)
	fmt.Println(`Also ensure it's not in an obvious place, like your desktop, hide it e.g. in C:\Program Files (x86)\Microsoft\EdgeCore`)
	fmt.Printf("Enter the path to save the config file: ")
	_, err := fmt.Scan(&newpath)
	if err != nil {
		log.Fatalf("%v", err)
		return
	}

	// open the config
	config, _ := os.Open(TempFiles + "\\HeadshotConfigs\\Baseinput.json")
	data, err := ioutil.ReadAll(config)
	if err != nil {
		log.Fatalf("%v", err)
		return
	}
	config.Close()

	// generate a random uuid
	uuid := uuid.New().String()

	// dump the config
	os.Create(newpath + "\\" + uuid + ".json")
	err = ioutil.WriteFile(newpath+"\\"+uuid+".json", data, 0)
	if err != nil {
		log.Fatalf("%v", err)
		return
	}
	log.Println("Config backed up to:", newpath+"\\"+uuid+".json")
}

func DeleteConfig() {
	homeDir := UserHomeDir()
	TempFiles := homeDir + "\\AppData\\Local\\Packages\\StudioWildcard.4558480580BB9_1w2mm55455e38\\AC\\Temp"

	// delete the config
	e := os.RemoveAll(TempFiles)
	os.Mkdir(TempFiles, 0755)
	if e != nil {
		log.Fatalf("%v", e)
		return
	}
	log.Println("Config deleted from:", TempFiles)
}

func UserHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}
