package cli

import (
	"fmt"
	"log"
	"time"

	"github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/plumbing/object"
	"github.com/iamsahebgiri/twig/internal/renderer"
)

type Log struct {
	Message string
	Author  string
	Date    time.Time
}

func Build(path string) {
	repo, err := git.PlainOpen(path)
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = repo.Close() }()

	head, err := repo.Head()
	if err != nil {
		log.Fatal(err)
	}

	r := renderer.New()

	commitIter, err := repo.Log(&git.LogOptions{From: head.Hash()})
	if err != nil {
		log.Fatal(err)
	}
	var logs []Log
	commitIter.ForEach(func(c *object.Commit) error {
		logs = append(logs, Log{
			Message: c.Message,
			Author:  c.Author.Name,
			Date:    time.Now(),
		})
		return nil
	})

	fmt.Println(len(logs))

	r.RenderFile("dist/index.html", "index", map[string]any{
		"Logs": logs,
	})
}
