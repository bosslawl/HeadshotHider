package hide

import (
	"HeadshotHider/core/backend"
	"HeadshotHider/core/util"
	"os"
)

func ClearCrashReports() error {
	reportsDir := backend.Directories["ReportsDir"]

	err := os.RemoveAll(reportsDir)
	if err != nil {
		util.Logger.Printf("Error clearing reports directory %s, error: %v\n", reportsDir, err)
		return err
	}
	util.Logger.Printf("Cleared reports directory %s\n", reportsDir)

	err = os.MkdirAll(reportsDir, 0755)
	if err != nil {
		util.Logger.Printf("Error creating reports directory %s, error: %v\n", reportsDir, err)
		return err
	}
	util.Logger.Printf("Recreated reports directory %s\n", reportsDir)

	return nil
}
