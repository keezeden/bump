package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
)

const usage = `
	Usage:
	  bump <package> <version> [repos]
	  bump <package> <version> -r <folder>
	  bump -h
	  
	Options:
	  -r	recursively bump subfolder repos

	Example:			
	  @keez/core_lib .27 cli_app web_app
		`

type Branch struct {
	Repo string
}

// Create new branch for version bump.
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

func main() {
	flag.Usage = func() {
		fmt.Print(usage)
	}

	var pkg, version, repos, recursive = parseArgs()

	if recursive {
		fmt.Println(recursive)
	}

	for _, repo := range repos {

		path := fmt.Sprintf("./%s", repo)

		branch := Branch{path}
		branch.New()

		file := readPkg(repo)

		var nFile, err = writePkg(file, pkg, version)

		if err != nil {
			log.Fatalf("Error occured while updating %s: %v", repo, err)
		}

		updatePkg(repo, nFile)
		fmt.Println("Updating Package...")

		branch.Commit()
		branch.Push()
	}
}

func parseArgs() (string, string, []string, bool) {
	recursive := flag.Bool("r", false, "")

	flag.Parse()
	args := flag.Args()

	pkg := args[0]
	version := args[1]
	repos := args[2:]

	return pkg, version, repos, *recursive
}

// Read repo package.json file into memory.
func readPkg(repo string) []string {
	file, err := ioutil.ReadFile(fmt.Sprintf("%s/package.json", repo))
	if err != nil {
		log.Fatalf("%v", err)
	}

	lines := strings.Split(string(file), "\n")

	return lines
}

// Write repo package.json in memory with new version.
func writePkg(lines []string, pkg string, version string) (string, error) {
	for i, line := range lines {
		if strings.Contains(line, pkg) {
			lines[i] = newVersion(line, version)
		}
	}

	return strings.Join(lines, "\n"), nil
}

// Write repo package.json in memory to file.
func updatePkg(repo string, nFile string) {
	err := ioutil.WriteFile(fmt.Sprintf("%s/package.json", repo), []byte(nFile), 0644)
	if err != nil {
		log.Fatalf("%v", err)
	}
}

// Format new version.
func newVersion(line string, newVersion string) string {
	// TODO . seperator for easy version updates
	split := strings.Split(line, ":")
	split[1] = fmt.Sprintf("\"%s\",", newVersion)
	fmt.Println("New Version", strings.Join(split, ":"))
	return strings.Join(split, ":")
}
