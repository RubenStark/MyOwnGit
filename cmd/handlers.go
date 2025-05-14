package cmd

import (
	"compress/zlib"
	"fmt"
	"io"
	"os"

	handlers "github.com/codecrafters-io/git-starter-go/handlers"
)

type Handler func([]string) string

var Handlers map[string]Handler = map[string]Handler{
	"init":        initHandler,
	"cat-file":    catFile,
	"hash-object": handlers.HashObject,
	"ls-tree":     handlers.LsTree,
	"write-tree":  handlers.WriteTree,
	"commit-tree": handlers.CommitTree,
}

func initHandler(args []string) string {

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

func catFile(args []string) string {
	if len(args) != 3 {
		fmt.Fprintf(os.Stderr, "usage: mygit cat-file <hash>\n")
		os.Exit(1)
	}

	r, err := zlib.NewReader(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating zlib reader: %s\n", err)
		return fmt.Sprintf("Error creating zlib reader: %s", err)
	}

	defer r.Close()

	var buf []byte

	if _, err := io.Copy(os.Stdout, r); err != nil {
		fmt.Fprintf(os.Stderr, "Error copying data: %s\n", err)
		return fmt.Sprintf("Error copying data: %s", err)
	}

	if err := r.Close(); err != nil {
		fmt.Fprintf(os.Stderr, "Error closing zlib reader: %s\n", err)
		return fmt.Sprintf("Error closing zlib reader: %s", err)
	}

	return fmt.Sprintf("File contents: %s", string(buf))
}
