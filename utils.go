package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
)

func folderExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}

	return info.IsDir()
}

func initializeRootSnippetsFolder(folderPath string) {

	if folderExists(folderPath) {
		fmt.Printf("The folder %s exists.\n", folderPath)
	} else {
		fmt.Printf("The folder %s does not exist.\n", folderPath)
		err := os.MkdirAll(folderPath, 0755)
		if err != nil {
			fmt.Println("Error creating folder")
			return
		}
		fmt.Println("Folder created successfully")
	}

}

func getHomeDir() (string, error) {
	if runtime.GOOS == "windows" {
		homeDir := os.Getenv("USERPROFILE")
		if homeDir == "" {
			return "", fmt.Errorf("USERPROFILE not set")
		}
		return homeDir, nil
	}

	homeDir := os.Getenv("HOME")
	if homeDir == "" {
		return "", fmt.Errorf("HOME not set")
	}
	return homeDir, nil
}
func executeCommand(command []string, location string) error {
	// Create the command

	cmd := exec.Command(command[0], command[1:]...)

	// Set the directory where the command should be executed
	cmd.Dir = location

	// Run the command and capture the output
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
		return err
	}

	// Print the output
	fmt.Println(string(output))
	return nil
}
