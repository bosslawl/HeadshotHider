package hide

import (
	"io"
	"os"
	"path/filepath"

	"HeadshotHider/core/backend"
	"HeadshotHider/core/util"

	"github.com/google/uuid"
)

func (c *Client) CopyConfig() error {
	config := backend.Directories["ConfigDir"]
	cpath, _ := os.ReadDir(config)

	for _, cfile := range cpath {
		if !cfile.IsDir() && filepath.Ext(cfile.Name()) == ".json" {
			source := filepath.Join(config, cfile.Name())
			destination := filepath.Join(c.DownloadPath, CfgName())
			aerr := CopyFile(source, destination)
			if aerr != nil {
				util.Logger.Println("Error copying file:", aerr)
				return aerr
			}
		}
	}
	return nil
}

func CfgName() string {
	uuid := uuid.New().String()

	return uuid + ".json"
}

func CopyFile(sourcePath, destinationPath string) error {
	sourceFile, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(destinationPath)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}

	return nil
}

func DeleteConfig() error {
	config := backend.Directories["ConfigDir"]
	os.RemoveAll(config)
	os.Mkdir(config, 0755)
	return nil
}

func (c *Client) ClearConfig() error {
	err := c.CopyConfig()
	if err != nil {
		util.Logger.Println("Error copying config:", err)
		return err
	}

	errr := DeleteConfig()
	if errr != nil {
		util.Logger.Println("Error deleting config:", errr)
		return errr
	}
	return nil
}
