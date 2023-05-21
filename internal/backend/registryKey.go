package backend

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/google/uuid"
	"golang.org/x/sys/windows/registry"
)

func (c *Client) CopyKey() {
	key, _ := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Headshot`, registry.QUERY_VALUE)
	value, _, err := key.GetStringValue("Key")
	if err != nil {
		panic(err)
	}

	uuid := uuid.New().String()
	os.Create(c.DownloadPath + "\\" + uuid + ".txt")
	ioutil.WriteFile(c.DownloadPath+"\\"+uuid+".txt", []byte(value), 0)
}

func DeleteTable() {
	key, _ := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Headshot`, registry.ALL_ACCESS)
	DeleteRegistryKey(key)
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

func (c *Client) RegistryKey() error {
	c.CopyKey()
	DeleteTable()
	return nil
}
