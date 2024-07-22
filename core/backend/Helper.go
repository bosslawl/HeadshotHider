package backend

import (
	"HeadshotHider/core/util"
	"errors"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func UserHomeDir() (string, error) {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		if home == "" {
			util.Logger.Println("Unable to determine user home directory")
			return "", errors.New("unable to determine user home directory")
		}
		return home, nil
	}
	home := os.Getenv("HOME")
	if home == "" {
		util.Logger.Println("Unable to determine user home directory")
		return "", errors.New("unable to determine user home directory")
	}
	return home, nil
}

func initDirectories() (map[string]string, error) {
	home, err := UserHomeDir()
	if err != nil {
		log.Printf("Error getting user home directory: %v", err)
		return nil, err
	}

	return map[string]string{
		"DownloadsDir": filepath.Join(home, "Downloads"),
		"BinDir":       filepath.Join(home, "AppData", "Local", "Temp", "bin_files"),
		"ConfigDir":    filepath.Join(home, "AppData", "Local", "Packages", "StudioWildcard.4558480580BB9_1w2mm55455e38", "AC", "Temp"),
		"CrashDir":     filepath.Join(home, "AppData", "Local", "CrashDumps"),
		"RecentDir":    filepath.Join(home, "AppData", "Roaming", "Microsoft", "Windows", "Recent"),
		"PrefetchDir":  "C:\\Windows\\Prefetch",
		"ReportsDir":   "C:\\ProgramData\\Microsoft\\Windows\\WER\\ReportArchive",
		"DefendDir":    "C:\\ProgramData\\Microsoft\\Windows Defender\\Scans\\History\\Service",
	}, nil
}

var Directories, _ = initDirectories()
