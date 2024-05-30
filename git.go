package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/v62/github"
)

func createRepo() {

	client := github.NewClient(nil).WithAuthToken()

	// create a new private repository named "foo"
	repo := &github.Repository{
		Name:    github.String("test_repo"),
		Private: github.Bool(true),
	}

	repo, _, err := client.Repositories.Create(context.Background(), "", repo)

	if err != nil {
		fmt.Println("Failed to create repo")
		return
	}

	fmt.Println("Repo created successfully")

}
