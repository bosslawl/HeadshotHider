package backend

import "os"

func DeleteHistory() {
	os.RemoveAll("C:\\ProgramData\\Microsoft\\Windows Defender\\Scans\\History\\Service")
	os.Mkdir("C:\\ProgramData\\Microsoft\\Windows Defender\\Scans\\History\\Service", 0755)
}