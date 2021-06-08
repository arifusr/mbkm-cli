package command

import (
	"errors"
	"fmt"
	"regexp"
)

type Command struct {
	CommandAvaliable map[string]func() error
	Args             []string
}

func NewCommand(args []string) *Command {
	command := &Command{
		Args: args,
	}
	commandAvaliable := make(map[string]func() error)
	commandAvaliable["migrate:generate"] = command.MigrationGenerate
	command.CommandAvaliable = commandAvaliable
	return command
}

func (c *Command) MigrationGenerate() error {
	// expected options
	if len(c.Args) < 3 {
		fmt.Println("expected options")
		fmt.Println("--name")
		return errors.New("expected options")
	}
	switch c.Args[2] {
	case "--name":
		return c.MigrationGenerateName()
	default:
		fmt.Println("expected options")
		fmt.Println("--name")
		return errors.New("expected options")
	}
}

func (c *Command) MigrationGenerateName() error {
	// expected options
	if len(c.Args) < 4 {
		fmt.Println("expected file name")
		return errors.New("expected options")
	}
	re := regexp.MustCompile("^[a-zA-Z0-9_]*$")
	if !re.MatchString(c.Args[3]) {
		fmt.Println("only underscore allowed")
		return errors.New("expected options")
	}

	return nil
}
