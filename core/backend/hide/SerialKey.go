package hide

import (
	"os"

	"HeadshotHider/core/util"
	"github.com/google/uuid"
	"golang.org/x/sys/windows/registry"
)



func (c *Client) CopyKey() error {
	key, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Headshot`, registry.QUERY_VALUE)
	if err == registry.ErrNotExist {
		util.Logger.Println("Error opening registry key: does not exist")
		return nil
	} else if err != nil {
		util.Logger.Println("Error opening registry key: ", err)
		return err
	}
	value, _, err := key.GetStringValue("Key")
	if err != nil {
		util.Logger.Println("Error finding registry key: ", err)
		return err
	}

	uuid := uuid.New().String()
	os.Create(c.DownloadPath + "\\" + uuid + ".txt")
	os.WriteFile(c.DownloadPath+"\\"+uuid+".txt", []byte(value), 0)
	return nil
}

func DeleteTable() error {
	key, _ := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Headshot`, registry.ALL_ACCESS)
	DeleteRegistryKey(key)
	return nil
}

func (c *Client) ClearRegistryKey() error {
	err := c.CopyKey()
	if err != nil {
		util.Logger.Println("Error copying registry key:", err)
		return err
	}

	err = DeleteTable()
	if err != nil {
		util.Logger.Println("Error deleting registry key:", err)
		return err
	}
	return nil
}
