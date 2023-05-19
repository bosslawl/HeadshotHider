package core

import (
	"fmt"
	"os"
	"path/filepath"
)

/*
	Deletes all files in in the root directory that are named "HSLoader.exe"
	Scans whole pc which is slow but it works (for now i won't be using it)
*/

func Loader() {
	searchDir := "C:\\"
	err := filepath.Walk(searchDir, func(path string, info os.FileInfo, err error) error {
		// Check if the current file matches the name "HSLoader.exe"
		if err == nil && info.Name() == "HSLoader.exe" {
			// Attempt to delete the file
			err = os.Remove(path)
			if err != nil {
				// Handle the error if the file cannot be deleted
				fmt.Println("Error deleting file:", err)
				return err
			}

			// Print a message indicating that the file was deleted
			fmt.Println("File deleted:", path)
		}
		return nil
	})

	if err != nil {
		// Handle the error if the search encounters an issue
		fmt.Println("Error searching for files:", err)
	}
}

func LoaderUpdated() {
	if _, err := os.Stat("HSLoader.exe"); os.IsNotExist(err) {
		fmt.Println("Please put this file in the same directory as HSLoader.exe")
	} else {
		// if file exists
		// delete file
		err := os.Remove("HSLoader.exe")
		if err != nil {
			// Handle the error if the file cannot be deleted
			fmt.Println("Error deleting file:", err)
		}
	}
}
