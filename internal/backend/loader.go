package backend

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func (c *Client) DeleteLoader() error {
	lpath, _ := ioutil.ReadDir(c.LoaderPath)

	for _, lfile := range lpath {
		if strings.Contains(strings.ToLower(lfile.Name()), "headshot") || strings.Contains(strings.ToLower(lfile.Name()), "hs") {
			os.RemoveAll(filepath.Join(c.LoaderPath, lfile.Name()))
		}
	}
	return nil
}
