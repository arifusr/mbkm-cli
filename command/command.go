package command

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/arifusr/mbkm-cli/file"
)

type Command struct {
	CommandAvaliable map[string]func() error
	Args             []string
	File             *file.File
}

func NewCommand(args []string) *Command {
	command := &Command{
		Args: args,
		File: file.NewFile(),
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

	// create file with signature of date

	filename := "filename"
	c.File.SetContent("aaaa")
	c.File.SetName(filename)
	c.File.SetDirPath("./")
	c.File.WriteFile()

	return nil
}