package hide

import (
	"os"
	"path/filepath"
	"strings"

	"HeadshotHider/core/util"

	"github.com/google/uuid"
)

func (c *Client) CopyLoader() error {
	loader := c.LoaderPath
	cpath, _ := os.ReadDir(loader)

	for _, cfile := range cpath {
		if !cfile.IsDir() && filepath.Ext(cfile.Name()) == ".exe" {
			source := filepath.Join(loader, cfile.Name())
			destination := filepath.Join(c.DownloadPath, LoaderName())
			aerr := CopyFile(source, destination)
			if aerr != nil {
				util.Logger.Println("Error copying file:", aerr)
				return aerr
			}
		}
	}
	return nil
}

func LoaderName() string {
	uuid := uuid.New().String()

	return uuid + ".exe"
}

func (c *Client) DeleteLoader() error {
	lpath, _ := os.ReadDir(c.LoaderPath)

	for _, lfile := range lpath {
		if strings.Contains(strings.ToLower(lfile.Name()), "headshot") || strings.Contains(strings.ToLower(lfile.Name()), "hs") {
			err := os.RemoveAll(filepath.Join(c.LoaderPath, lfile.Name()))
			if err != nil {
				util.Logger.Println("Error deleting loader:", err)
				return err
			}
		}
	}
	return nil
}

func (c *Client) ClearLoader() error {
	err := c.CopyLoader()
	if err != nil {
		util.Logger.Println("Error copying loader:", err)
		return err
	}

	err = c.DeleteLoader()
	if err != nil {
		util.Logger.Println("Error deleting loader:", err)
		return err
	}
	return nil
}
