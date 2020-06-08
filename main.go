package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const usage = `
  Bump, the multi-repo package version bumping tool.

	Usage:
			
	  @author/package version repo1 repo2 repo3

	Example:
				
	  @keez/core_lib .27 cli_app web_app
		`

// Initializing overrides.
func setup() {
	flag.Usage = func() {
		fmt.Print(usage)
	}
}

func main() {
	setup()

	var pkg, version, repos, recursive = parseArgs()

	fmt.Println(recursive)

	for _, repo := range repos {
		var file = readPkg(repo)
		var nFile, err = writePkg(file, repo, pkg, version)
		if err != nil {
			log.Fatalln("Error occured while updating %s: %s", repo, err)
		}
		updatePkg(repo, nFile)
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

func readPkg(repo string) []string {
	file, err := ioutil.ReadFile(fmt.Sprintf("%s/package.json", repo))
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(file), "\n")

	return lines
}

func writePkg(lines []string, repo string, pkg string, version string) (string, error) {
	for i, line := range lines {
		if strings.Contains(line, pkg) {
			lines[i] = newVersion(line, version)
		}
	}

	return strings.Join(lines, "\n"), nil
}

func updatePkg(repo string, nFile string) {
	err := ioutil.WriteFile(fmt.Sprintf("%s/package_test.json", repo), []byte(nFile), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}

func newVersion(line string, newVersion string) string {
	split := strings.Split(line, ":")
	split[1] = newVersion
	return strings.Join(split, ":")
}
