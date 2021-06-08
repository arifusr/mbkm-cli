package main

import (
	"os"

	"github.com/arifusr/mbkm-cli/command"
	"github.com/arifusr/mbkm-cli/validation"
)

func main() {
	args := os.Args
	cmd := command.NewCommand(args)

	validator := validation.NewValidator(cmd)
	if err := validator.ValidateCommand(args); err != nil {
		return
	}
	//run command
	cmd.CommandAvaliable[args[1]]()
}
