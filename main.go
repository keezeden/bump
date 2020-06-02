package main

import (
	"flag"
	"fmt"
)

func setup() {
	// forgive me for i have sinned
	flag.Usage = func() {
		fmt.Println("\n")
		fmt.Println("Bump, the multi-repo package version bumping tool.")
		fmt.Println("\n")
		fmt.Println("Usage:")
		fmt.Println("\n")
		fmt.Println("	@author/package version repo1 repo2 repo3")
		fmt.Println("\n")
		fmt.Println("Example:")
		fmt.Println("\n")
		fmt.Println("	@keez/core_lib .27 cli_app web_app \n")
		fmt.Println("\n")
		flag.PrintDefaults()
	}
}

func main() {

	setup()

	recursive := flag.Bool("r", false, "Recursively bump all child repos in selected folder.")
	path := flag.String("p")

	flag.Parse()

	args := flag.Args()

	module := args[0]
	version := args[1]
	repos := args[2:]

	fmt.Println("Module to bump:", module)
	fmt.Println("Version number:", version)
	fmt.Println("Repos to update:", repos)
	fmt.Println("Recursive?:", *recursive)

}
