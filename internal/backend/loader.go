package backend

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func (c *Client) DeleteLoader() {
	files, _ := ioutil.ReadDir(c.LoaderPath)

	for _, file := range files {
		if strings.Contains(strings.ToLower(file.Name()), "headshot") || strings.Contains(strings.ToLower(file.Name()), "hs") {
			os.Remove(filepath.Join(c.LoaderPath, file.Name()))
		}
	}
}
