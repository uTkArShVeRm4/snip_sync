package main

import (
	"fmt"
	"os"

	"github.com/lithammer/fuzzysearch/fuzzy"
)

type Snippet struct {
	Name string
	Buf  string
	Rank int
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
