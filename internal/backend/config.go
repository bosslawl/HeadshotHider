package backend

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/bosslawl/HeadshotHider/v2/internal/util"
	"github.com/google/uuid"
)

func (c *Client) CopyConfig() error {
	config := UserHomeDir() + "\\AppData\\Local\\Packages\\StudioWildcard.4558480580BB9_1w2mm55455e38\\AC\\Temp\\HeadshotConfigs"
	cpath, _ := ioutil.ReadDir(config)

	for _, cfile := range cpath {
		if !cfile.IsDir() && filepath.Ext(cfile.Name()) == ".json" {
			source := filepath.Join(config, cfile.Name())
			destination := filepath.Join(c.DownloadPath, NewName())
			CopyFile(source, destination)
		}
	}
	return nil
}

func NewName() string {
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
	config := UserHomeDir() + "\\AppData\\Local\\Packages\\StudioWildcard.4558480580BB9_1w2mm55455e38\\AC\\Temp"
	os.RemoveAll(config)
	os.Mkdir(config, 0755)
	return nil
}

func (c *Client) Config() error {
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
