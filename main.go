package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/lithammer/fuzzysearch/fuzzy"
)

type Snippet struct {
	Name string
	Buf  string
	Rank int
}

func folderExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}

	return info.IsDir()
}

func initializeSnipFolder(folderPath string) {

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

func readSnippets(folderPath string) []Snippet {
	var snippets []Snippet
	entries, err := os.ReadDir(folderPath)
	if err != nil {
		fmt.Println(err)
		return snippets
	}
	for _, file := range entries {
		if file.IsDir() {
			continue
		}
		path := folderPath + "/" + file.Name()
		file, err := os.Open(path)
		if err != nil {
			fmt.Println(err)
			continue
		}

		buf := make([]byte, 128)
		n, err := file.Read(buf)
		if err != nil {
			file.Close()
			fmt.Println(err)
			continue
		}
		file.Close()

		if n < 128 {
			buf = buf[:n]
		}

		snippet := Snippet{
			Name: file.Name(),
			Buf:  string(buf),
			Rank: 0,
		}
		snippets = append(snippets, snippet)
	}
	return snippets
}

func (s *Snippet) CalculateRank(target string) {
	s.Rank = fuzzy.RankMatch(target, s.Buf)
}

func main() {
	snippetFolder := "./snippets"
	initializeSnipFolder(snippetFolder)
	snippets := readSnippets(snippetFolder)

	target := "tis"

	for i := range snippets {
		snippets[i].CalculateRank(target)
	}

	sort.Slice(snippets, func(i, j int) bool {
		return snippets[i].Rank > snippets[j].Rank
	})

	for _, snippet := range snippets {
		fmt.Println(snippet.Name, snippet.Rank)
	}
}
