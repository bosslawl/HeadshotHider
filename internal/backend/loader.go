package backend

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/bosslawl/HeadshotHider/v2/internal/util"
)

func (c *Client) DeleteLoader() error {
	lpath, _ := ioutil.ReadDir(c.LoaderPath)

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
