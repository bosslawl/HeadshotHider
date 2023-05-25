package backend

import (
	"os"
	"path/filepath"

	"github.com/bosslawl/HeadshotHider/v2/internal/util"
)

func ClearCrashReports() error {
	filepath.Walk("C:\\ProgramData\\Microsoft\\Windows\\WER\\ReportArchive", func(path string, info os.FileInfo, err error) error {
		rem := os.RemoveAll(path)
		if rem != nil {
			util.Logger.Println("Error deleting Windows Error Reporting:", rem)
			return rem
		}
		mk := os.Mkdir("C:\\ProgramData\\Microsoft\\Windows\\WER\\ReportArchive", 0755)
		if mk != nil {
			util.Logger.Println("Error creating Windows Error Reporting directory:", mk)
			return mk
		}
		return nil
	})
	return nil
}
