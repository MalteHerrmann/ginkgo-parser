// ginkgo-parser is a tool to parse the contents of a Go test file with BDD style testing
// using the GinkGo framework.
//
// It prints the specification to the terminal, when executed using
//
//	go run github.com/MalteHerrmann/ginkgo-parser [FILEPATH]
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

func main() {
	// Get the filename from command-line argument
	if len(os.Args) < 2 {
		log.Fatalf("Usage: go run filename.go <filename>")
	}
	filename := os.Args[1]

	// Read the Go file
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}

	printBlocks(string(content))
}

func printBlocks(blocks string) {
	// Define unified regex pattern to match Describe, Context, and It blocks
	regexPattern := `([ \t]*)(Describe|Context|It)\(\s*"(.*?)"\s*,\s*func\(\)\s*{([\s\S]*?)\}\s*\)`

	re := regexp.MustCompile(regexPattern)
	matches := re.FindAllStringSubmatch(blocks, -1)

	for _, match := range matches {
		whiteSpace := match[1]
		blockName := match[3]

		fmt.Printf("%s- %s\n", whiteSpace, blockName)
	}
}
