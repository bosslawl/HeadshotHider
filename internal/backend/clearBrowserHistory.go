package backend

import (
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/bosslawl/HeadshotHider/v2/internal/util"
)

const (
	chrome      = "chrome.exe"
	opera       = "opera.exe"
	operaGX     = "Opera.exe"
	firefox     = "firefox.exe"
	edge        = "msedge.exe"
	brave       = "brave.exe"
	vivaldi     = "vivaldi.exe"
	yandex      = "yandex.exe"
	tor         = "tor.exe"
	chromium    = "chromium.exe"
	iridium     = "iridium.exe"
	epic        = "epicbrowser.exe"
	comodo      = "comodo.exe"
	srware      = "srware.exe"
	torch       = "torch.exe"
	duckduckgo  = "duckduckgo.exe"
	centBrowser = "centbrowser.exe"
	slimjet     = "slimjet.exe"
	maxthon     = "maxthon.exe"
)

var browsers = map[string]string{
	chrome:      "\\AppData\\Local\\Google\\Chrome\\User Data\\Default",
	opera:       "\\AppData\\Roaming\\Opera Software\\Opera Stable",
	operaGX:     "\\AppData\\Roaming\\Opera Software\\Opera GX Stable",
	firefox:     "\\AppData\\Roaming\\Mozilla\\Firefox\\Profiles",
	edge:        "\\AppData\\Local\\Microsoft\\Edge\\User Data\\Default",
	brave:       "\\AppData\\Local\\BraveSoftware\\Brave-Browser\\User Data\\Default",
	vivaldi:     "\\AppData\\Local\\Vivaldi\\User Data\\Default",
	yandex:      "\\AppData\\Local\\Yandex\\YandexBrowser\\User Data\\Default",
	tor:         "\\AppData\\Local\\Tor Browser\\Browser\\TorBrowser\\Data\\Browser\\profile.default",
	chromium:    "\\AppData\\Local\\Chromium\\User Data\\Default",
	iridium:     "\\AppData\\Local\\Iridium\\User Data\\Default",
	epic:        "\\AppData\\Local\\Epic Privacy Browser\\User Data\\Default",
	comodo:      "\\AppData\\Local\\Comodo\\Dragon\\User Data\\Default",
	srware:      "\\AppData\\Local\\SRWare Iron\\User Data\\Default",
	torch:       "\\AppData\\Local\\Torch\\User Data\\Default",
	duckduckgo:  "\\AppData\\Local\\DuckDuckGo\\User Data\\Default",
	centBrowser: "\\AppData\\Local\\CentBrowser\\User Data\\Default",
	slimjet:     "\\AppData\\Local\\Slimjet\\User Data\\Default",
	maxthon:     "\\AppData\\Local\\Maxthon\\User Data\\Default",
}

func deleteBrowserHistory(browserExe, browserDir string) error {
	exec.Command("taskkill", "/F", "/IM", browserExe).Run()
	time.Sleep(5 * time.Second)
	err := os.Remove(filepath.Join(UserHomeDir() + browserDir, "History"))
	if err != nil {
		util.Logger.Println("Error deleting browser history:", err)
		return err
	}
	return nil
}

func ClearBrowserHistory() error {
	for browserExe, browserDir := range browsers {
		if _, err := os.Stat(UserHomeDir() + browserDir); err == nil {
			deleteBrowserHistory(browserExe, browserDir)
			return nil
		}
	}
	return nil
}