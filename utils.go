package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func parseJsonFilesInDirectory(pathname string) {
	entries, err := os.ReadDir(pathname)

	if err != nil {
		fmt.Println("Error reading directory", err)
		return
	}

	for _, entry := range entries {
		resultingPath := filepath.Join(pathname, entry.Name())

		if entry.IsDir() {
			fmt.Printf("Directory:\n%s\n\n", resultingPath)
			parseJsonFilesInDirectory(resultingPath)
		} else {
			parseJsonFile(resultingPath)
			fmt.Println()
		}
	}
}

func parseJsonFile(pathname string) {
	data, err := os.ReadFile(pathname)

	if err != nil {
		fmt.Println("Error reading input", err)
		return
	}

	tokens := lexer(string(data))
	result := parser(tokens)

	fmt.Printf("File:\n%s\n\nInput:\n%s\n\nResult:\n%s\n", pathname, string(data), result)
}
