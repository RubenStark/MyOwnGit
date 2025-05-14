package handlers

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"fmt"
	"os"
)

func HashObject(args []string) string {
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
