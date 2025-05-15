package handlers

import (
	"fmt"
	"os"
)

func InitHandler(args []string) string {

	fmt.Println("Initializing empty repository...")

	for _, dir := range []string{".git", ".git/objects", ".git/refs"} {
		if err := os.MkdirAll(dir, 0755); err != nil {
			fmt.Fprintf(os.Stderr, "Error creating directory: %s\n", err)
			return fmt.Sprintf("Error creating directory: %s", err)
		}
	}

	headFileContents := []byte("ref: refs/heads/main\n")
	if err := os.WriteFile(".git/HEAD", headFileContents, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing file: %s\n", err)
		return "Error writing .git/HEAD"
	}

	return "Initialized empty repository"
}
