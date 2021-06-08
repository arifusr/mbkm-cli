package main

import (
	"fmt"
	"os"

	"github.com/arifusr/mbkm-cli/command"
	"github.com/arifusr/mbkm-cli/validation"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Print("error load env")
		return
	}
	args := os.Args
	cmd := command.NewCommand(args)

	validator := validation.NewValidator(cmd)
	if err := validator.ValidateCommand(args); err != nil {
		return
	}
	//run command
	cmd.CommandAvaliable[args[1]]()
}
