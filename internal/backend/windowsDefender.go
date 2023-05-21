package backend

import "os"

func DeleteHistory() error {
	os.RemoveAll("C:\\ProgramData\\Microsoft\\Windows Defender\\Scans\\History\\Service")
	os.Mkdir("C:\\ProgramData\\Microsoft\\Windows Defender\\Scans\\History\\Service", 0755)
	return nil
}