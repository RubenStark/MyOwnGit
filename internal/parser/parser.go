package parser

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/codecrafters-io/git-starter-go/internal/command"
)

type Parser struct {
	Reader *bufio.Reader
}

func New(r *bufio.Reader) Parser {
	return Parser{
		Reader: r,
	}
}

func (p Parser) ParseInput() (command.Command, error) {

	line, err := p.Reader.ReadString('\n')
	if err != nil {
		return command.Command{}, err
	}

	line = strings.TrimSpace(line)
	args := strings.Fields(line)

	if len(args) < 1 {
		return command.Command{}, fmt.Errorf("not enough arguments")
	}

	return command.Command{
		Name: strings.ToLower(args[0]),
		Args: args[1:],
	}, nil

}
