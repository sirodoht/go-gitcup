package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"gopkg.in/src-d/go-git.v4"
)

func main() {
	filePtr := flag.String("f", "-", "File with list of repo URL to backup")
	flag.Parse()

	if *filePtr != "-" {
		readFromFile(*filePtr)
	} else {
		repoURL := os.Args[1]
		fmt.Println("Git repo URL:", repoURL)
		processRepo(repoURL)
	}
}

func readFromFile(filePath string) {
	fmt.Println("Reading repo file:", filePath)

	file, err := os.Open(filePath)
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	// Start reading from the file with a reader.
	reader := bufio.NewReader(file)

	var line string
	for {
		line, err = reader.ReadString('\n')

		repoLine := strings.Trim(line, "\n")
		processRepo(repoLine)

		if err != nil {
			break
		}
	}
}

func processRepo(repoURL string) {
	repoPath := "./repo-data/" + strings.Split(repoURL, "/")[1]
	if !strings.Contains(repoURL, "@") {
		repoPath = "./repo-data/" + strings.Split(repoURL, "/")[4]
	}
	fmt.Println("\nProcessing", repoPath)

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
			log.Println("Error:", err)
		}

		// Pull the latest changes from the origin remote and merge into the current branch
		err = worktree.Pull(&git.PullOptions{RemoteName: "origin"})
		if err != nil {
			if err.Error() == "already up-to-date" {
				log.Println("up to date")
			} else {
				log.Fatalln("Error", err)
			}
		} else {
			log.Println("pull success")

		}
	} else {
		// This case is for repo not existing, so we git clone
		_, err := git.PlainClone(repoPath, false, &git.CloneOptions{
			URL:      repoURL,
			Progress: os.Stdout,
		})

		if err != nil {
			log.Println("Error:", err)
		} else {
			log.Println("clone success")
		}
	}
}
