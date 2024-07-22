package hide

import (
	"HeadshotHider/core/backend"
	"HeadshotHider/core/util"
	"os"
	"path/filepath"
)

var browsers = map[string]string{
	"chrome.exe":  "\\AppData\\Local\\Google\\Chrome\\User Data\\Default",
	"firefox.exe": "\\AppData\\Roaming\\Mozilla\\Firefox\\Profiles",
	"msedge.exe":  "\\AppData\\Local\\Microsoft\\Edge\\User Data\\Default",
	"opera.exe":   "\\AppData\\Roaming\\Opera Software\\Opera Stable",
	"operagx.exe": "\\AppData\\Roaming\\Opera Software\\Opera GX Stable",
	"brave.exe":   "\\AppData\\Local\\BraveSoftware\\Brave-Browser\\User Data\\Default",
}

func ClearBrowsers(exe, dir string) error {
	homeDir, err := backend.UserHomeDir()
	if err != nil {
		util.Logger.Printf("Error getting user home directory: %v", err)
		return err
	}

	historyPath := filepath.Join(homeDir+dir, "History")
	if err := os.Remove(historyPath); err != nil {
		util.Logger.Printf("Error deleting browser history at %s, error: %v\n", historyPath, err)
		return err
	}

	util.Logger.Printf("Cleared browser history at %s\n", historyPath)
	return nil
}

func ClearHistory() error {
	for exe, dir := range browsers {
		homeDir, err := backend.UserHomeDir()
		if err != nil {
			util.Logger.Printf("Error getting user home directory: %v", err)
			return err
		}

		if _, err := os.Stat(homeDir + dir); err == nil {
			if err := ClearBrowsers(exe, dir); err != nil {
				util.Logger.Printf("Error clearing history for browser %s, error: %v\n", exe, err)
				return err
			}
		}
	}
	return nil
}
