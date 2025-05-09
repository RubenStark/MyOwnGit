package main

import (
	"fmt"
	"os"

	h "github.com/codecrafters-io/git-starter-go/cmd"
	"github.com/codecrafters-io/git-starter-go/internal/command"
)

// Usage: your_program.sh <command> <arg1> <arg2> ...
func main() {

	// We parse the input and get the command based on the Command struct
	var command = command.Command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	handler, ok := h.Handlers[command.Name]

	if !ok {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}

	output := handler(os.Args[1:])
	fmt.Fprintf(os.Stderr, output)

}
