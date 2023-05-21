package backend

import (
	"os"

	"github.com/bosslawl/HeadshotHider/v2/internal/util"
)

func DeleteHistory() error {
	err := os.RemoveAll("C:\\ProgramData\\Microsoft\\Windows Defender\\Scans\\History\\Service")
	if err != nil {
		util.Logger.Println("Error deleting Windows Defender history:", err)
		return err
	}
	os.Mkdir("C:\\ProgramData\\Microsoft\\Windows Defender\\Scans\\History\\Service", 0755)
	return nil
}
