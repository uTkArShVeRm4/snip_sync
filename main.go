package main

import (
	"fmt"
	"log"
	"path"
	"strings"
)

func main() {
	snippetFolder := "snippets"

	homedir, err := getHomeDir()
	if err != nil {
		log.Fatal(err)
		return
	}
	snippetsPath := path.Join(homedir, snippetFolder)
	initializeRootSnippetsFolder(snippetsPath)

	exists, token := checkGitCredentials()
	if !exists {
		fmt.Println("Git credentials not found")
		// ask user to enter auth token
		// save token to file
		fmt.Println("Enter your GitHub auth token")
		var token string
		fmt.Scanln(&token)
		saveGitCredentials(token)
	}

	executeCommand(strings.Fields("export git-credentials"), snippetsPath)

}
