package hide

import (
	"os"
	"os/exec"
	"strconv"
	"strings"

	"HeadshotHider/core/util"
)

func CloseARK() error {
	executableName := "shootergame.exe"

	cmd := exec.Command("tasklist", "/fo", "csv", "/nh")
	output, err := cmd.Output()
	if err != nil {
		util.Logger.Println("Error getting process list:", err)
		return err
	}

	pid, err := findProcessID(string(output), executableName)
	if err != nil {
		util.Logger.Println("Error getting ark process:", err)
		return err
	}

	if pid != 0 {
		err = killProcess(pid)
		if err != nil {
			util.Logger.Println("Error killing ark process:", err)
		}
	} else {
		util.Logger.Println("ARK not running")
		return nil
	}
	return nil
}

func findProcessID(output, executableName string) (int, error) {
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		fields := strings.Split(line, ",")
		if len(fields) > 1 {
			processName := strings.Trim(fields[0], `"`)
			processID, err := strconv.Atoi(strings.Trim(fields[1], `"`))
			if err != nil {
				return 0, err
			}

			if strings.EqualFold(processName, executableName) {
				return processID, nil
			}
		}
	}

	return 0, nil
}

func killProcess(pid int) error {
	process, err := os.FindProcess(pid)
	if err != nil {
		return err
	}

	err = process.Kill()
	if err != nil {
		return err
	}

	return nil
}
