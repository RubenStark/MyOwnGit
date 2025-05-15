package handlers

import (
	"os"
)

func Clone(args []string) string {

	repo := args[0]
	directory := args[1]

	// Create the directory if it doesn't exist
	err := os.MkdirAll(directory, 0755)
	if err != nil {
		return "Error creating directory: " + err.Error()
	}

	// and change to that directory
	err = os.Chdir(directory)
	if err != nil {
		return "Error changing directory: " + err.Error()
	}

	// Init the repo
	InitHandler(args)

	// Get packfile

	return "Cloned repository " + repo + " into " + directory

}
