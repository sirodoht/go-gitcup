package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"gopkg.in/src-d/go-git.v4"
)

func main() {
	repoULR := os.Args[1]
	fmt.Println("Git repo URL:", repoULR)

	repoPath := "./repo-data/" + strings.Split(repoULR, "/")[1]
	if !strings.Contains(repoULR, "@") {
		repoPath = "./repo-data/" + strings.Split(repoULR, "/")[4]
	}
	fmt.Println("Saved in ", repoPath)

	repoPathDotGit := repoPath + "/.git"
	// If repo exists, git pull
	if _, err := os.Stat(repoPathDotGit); !os.IsNotExist(err) {
		// We instantiate a new repository targeting the given path (the .git folder)
		repo, err := git.PlainOpen(repoPath)
		if err != nil {
			log.Fatal(err)
		}

		// Get working directory for the repository
		worktree, err := repo.Worktree()
		if err != nil {
			log.Fatal(err)
		}

		// Pull the latest changes from the origin remote and merge into the current branch
		err = worktree.Pull(&git.PullOptions{RemoteName: "origin"})
		if err != nil {
			log.Fatal(err)
		} else {
			dt := time.Now()
			fmt.Println(dt.Format("01-02-2006 15:04:05"), "pull success")
		}
	} else {
		// This case is for repo not existing, so we git clone
		_, err := git.PlainClone(repoPath, false, &git.CloneOptions{
			URL:      repoULR,
			Progress: os.Stdout,
		})

		if err != nil {
			log.Fatal(err)
		} else {
			dt := time.Now()
			fmt.Println(dt.Format("01-02-2006 15:04:05"), "clone success")
		}
	}
}
