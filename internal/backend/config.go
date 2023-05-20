package backend

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

func (c *Client) CopyConfig() {
	config := UserHomeDir() + "\\AppData\\Local\\Packages\\StudioWildcard.4558480580BB9_1w2mm55455e38\\AC\\Temp\\HeadshotConfigs"
	files, _ := ioutil.ReadDir(config)

	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".json" {
			source := filepath.Join(config, file.Name())
			destination := filepath.Join(c.DownloadPath, NewName())
			CopyFile(source, destination)
		}
	}
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

func DeleteConfig() {
	config := UserHomeDir() + "\\AppData\\Local\\Packages\\StudioWildcard.4558480580BB9_1w2mm55455e38\\AC\\Temp\\HeadshotConfigs"
	os.RemoveAll(config)
	os.Mkdir(config, 0755)
}

func (c *Client) Config() {
	c.CopyConfig()
	DeleteConfig()
}