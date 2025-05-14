package cmd

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"fmt"
	"io"
	"os"

	handlers "github.com/codecrafters-io/git-starter-go/handlers"
)

type Handler func([]string) string

var Handlers map[string]Handler = map[string]Handler{
	"init":        initHandler,
	"cat-file":    catFile,
	"hash-object": hashObject,
	"ls-tree":     handlers.LsTree,
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

func hashObject(args []string) string {
	if len(args) != 3 {
		fmt.Fprintf(os.Stderr, "usage: mygit hash-object <file>\n")
		os.Exit(1)
	}

	fileName := args[2]

	// Read the file contents
	fileContents, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %s\n", err)
		return fmt.Sprintf("Error reading file: %s", err)
	}

	// Create the Git object header
	objectHeader := fmt.Sprintf("blob %d\x00", len(fileContents))
	fullObject := append([]byte(objectHeader), fileContents...)

	// Compute the SHA-1 hash
	hash := sha1.Sum(fullObject)
	hashString := fmt.Sprintf("%x", hash)

	// Create the .git/objects/<first two characters>/<remaining characters> path
	objectDir := fmt.Sprintf(".git/objects/%s", hashString[:2])
	objectPath := fmt.Sprintf("%s/%s", objectDir, hashString[2:])

	// Ensure the directory exists
	if err := os.MkdirAll(objectDir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating object directory: %s\n", err)
		return fmt.Sprintf("Error creating object directory: %s", err)
	}

	// Compress the object and write it to the file
	var compressedObject bytes.Buffer
	w := zlib.NewWriter(&compressedObject)
	if _, err := w.Write(fullObject); err != nil {
		fmt.Fprintf(os.Stderr, "Error compressing object: %s\n", err)
		return fmt.Sprintf("Error compressing object: %s", err)
	}
	w.Close()

	if err := os.WriteFile(objectPath, compressedObject.Bytes(), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing object file: %s\n", err)
		return fmt.Sprintf("Error writing object file: %s", err)
	}

	return hashString
}
