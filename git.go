package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
)

type RepoExistsError struct{}

func (e *RepoExistsError) Error() string {
	return "Repo already exists"
}

func checkGitCredentials() (bool, string) {
	// read file at snippets/git-credentials
	homeDir, err := getHomeDir()
	file, err := os.Open(path.Join(homeDir, "snippets", "git-credentials"))
	if err != nil {
		fmt.Println("Error opening file")
		return false, ""
	}
	// read file, it must be in format export GITHUB_TOKEN=<token>
	// return true if the file contains the token and token
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "GITHUB_TOKEN=") {
			split := strings.Split(line, "=")
			if len(split[1]) > 0 {
				return true, split[1]
			}
		}
	}
	err = scanner.Err()
	if err != nil {
		fmt.Println("Error reading file")
		return false, ""
	}
	return false, ""
}

func saveGitCredentials(token string) {
	homeDir, err := getHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory")
		return
	}
	file, err := os.Create(path.Join(homeDir, "snippets", "git-credentials"))
	if err != nil {
		fmt.Println("Error creating file")
		return
	}
	defer file.Close()
	_, err = file.WriteString(fmt.Sprintf("GITHUB_TOKEN=%s", token))
	if err != nil {
		fmt.Println("Error writing to file")
		return
	}
}

func initializeNewSnippetRepo(rootFolderPath string, repoName string) error {
	folderPath := path.Join(rootFolderPath, repoName)
	if folderExists(folderPath) {
		return &RepoExistsError{}
	} else {
		fmt.Printf("The folder %s does not exist.\n", folderPath)
		err := os.MkdirAll(folderPath, 0755)
		if err != nil {
			fmt.Println("Error creating folder")
			return err
		}
		fmt.Println("Folder created successfully")
		_ = executeCommand(strings.Fields("git init"), folderPath)
		_ = executeCommand(strings.Fields("git checkout -b main"), folderPath)
		return nil
	}
}
func addCommitPush(repoPath string) {
	_ = executeCommand(strings.Fields("git add ."), repoPath)
	_ = executeCommand(strings.Fields("git commit -m 'Commit'"), repoPath)
	_ = executeCommand(strings.Fields("git push origin main"), repoPath)

}

//
// import (
// 	"context"
// 	"fmt"
//
// 	"github.com/google/go-github/v62/github"
// )
//
// func createRepo() {
//
//
// 	// create a new private repository named "foo"
// 	repo := &github.Repository{
// 		Name:    github.String("test_repo"),
// 		Private: github.Bool(true),
// 	}
//
// 	repo, _, err := client.Repositories.Create(context.Background(), "", repo)
//
// 	if err != nil {
// 		fmt.Println("Failed to create repo")
// 		return
// 	}
//
// 	fmt.Println("Repo created successfully")
//
// }
