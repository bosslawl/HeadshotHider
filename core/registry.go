package core

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/google/uuid"
	"golang.org/x/sys/windows/registry"
)

/*
	Start off by collecting the HS key from the registry
	Then dumps it to a random uuid file in the path the user specifies
	Then deletes the key from the registry
*/

func CopyRegistry() {
	key, _ := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Headshot`, registry.QUERY_VALUE)

	value, _, err := key.GetStringValue("Key")
	if err != nil {
		log.Fatalf("%v", err)
		return
	}

	var keypath string

	fmt.Println(`Choose where to save your key to so you don't loose it.`)
	fmt.Println(`This is your login key, so make sure you save it in a safe place (I recommend the same place as your config).`)
	fmt.Println(`Also ensure it's not in an obvious place, like your desktop, hide it e.g. in C:\Program Files (x86)\Microsoft\EdgeCore`)
	fmt.Printf("Enter the path to save the key file: ")
	_, err = fmt.Scan(&keypath)
	if err != nil {
		log.Fatalf("%v", err)
		return
	}

	// generate a random uuid
	uuid := uuid.New().String()

	// dump the config
	fmt.Println("1")
	os.Create(keypath + "\\" + uuid + ".txt")
	err = ioutil.WriteFile(keypath+"\\"+uuid+".txt", []byte(value), 0)
	if err != nil {
		log.Fatalf("%v", err)
		return
	}
	log.Println("Key backed up to:", keypath+"\\"+uuid+".txt")
}

func DeleteRegistry() {
	key, _ := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Headshot`, registry.ALL_ACCESS)

	DeleteRegistryKey(key)
	log.Println("Registry key deleted from: Computer\\HKEY_LOCAL_MACHINE\\SOFTWARE\\Headshot")
}

func DeleteRegistryKey(key registry.Key) error {
	// Find all the value names inside the key and delete them.
	valueNames, err := key.ReadValueNames(1000)
	if err != nil && err != io.EOF {
		return err
	}

	for _, valueName := range valueNames {
		err = key.DeleteValue(valueName)
		if err != nil && err != io.EOF {
			return err
		}
	}

	// Find the subkeys and delete those recursively.
	subkeyNames, err := key.ReadSubKeyNames(1000)
	if err != nil && err != io.EOF {
		return err
	}

	for _, subkeyName := range subkeyNames {
		subKey, err := registry.OpenKey(key, subkeyName, registry.ALL_ACCESS)
		if err != nil {
			return err
		}

		err = DeleteRegistryKey(subKey)
		if err != nil {
			return err
		}
	}

	// Then delete itself.
	return registry.DeleteKey(key, "")
}
