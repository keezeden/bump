package main

import (
	"os/exec"
	"log"
)

// Branch is...
type Branch struct {
	Repo string
}

// New branch for version bump.
func (b *Branch) New() {
	// Check if branch already exists.
	oldBranch := exec.Command("git", "checkout", "build/version-bump")
	oldBranch.Dir = b.Repo
	oErr := oldBranch.Run()
	if oErr != nil {
		b.Stash()

		branch := exec.Command("git", "checkout", "-b", "build/version-bump")
		branch.Dir = b.Repo
		bErr := branch.Run()
		if bErr != nil {
			log.Fatalf("an error occured while creating branch: %v", bErr)
		}
	}
}

// Stash any unsaved changes.
func (b *Branch) Stash() {
	stash := exec.Command("git", "stash", "-u", "-k")
	stash.Dir = b.Repo
	sErr := stash.Run()
	if sErr != nil {
		log.Fatalf("an error occured while stashing changes: %v", sErr)
	}
}

// Commit to branch.
func (b *Branch) Commit() {
	stage := exec.Command("git", "add", "--all", ".")
	stage.Dir = b.Repo
	commit := exec.Command("git", "commit", "-m", "build(bump): version bump")
	commit.Dir = b.Repo
	sErr := stage.Run()
	if sErr != nil {
		log.Fatalf("an error occured while staging changes: %v", sErr)
	}
	cErr := commit.Run()
	if cErr != nil {
		log.Fatalf("an error occured while comitting changes: %v", cErr)
	}
}

// Push changes to origin.
func (b *Branch) Push() {
	push := exec.Command("git", "push", "--set-upstream", "origin", "build/version-bump")
	push.Dir = b.Repo
	err := push.Run()
	if err != nil {
		log.Fatalf("an error occured while pushing changes: %v", err)
	}
}
