package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: json-parser -r <filename> | json-parser -d <directory>")
		os.Exit(1)
	}

	option := os.Args[1]
	pathname := os.Args[2]

	extension := filepath.Ext(pathname)

	switch option {
	case "-r":
		if extension == "" {
			fmt.Println("Error reading file, expected a file with an extension")
			os.Exit(1)
		}
	case "-d":
		if extension != "" {
			fmt.Println("Error reading directory, expected a directory path without a file extension")
			os.Exit(1)
		}
	}

	switch option {
	case "-r":
		parseJsonFile(pathname)
	case "-d":
		parseJsonFilesInDirectory(pathname)
	default:
		fmt.Printf("Invalid command: %s\n", option)
		os.Exit(1)
	}

	os.Exit(0)
}
