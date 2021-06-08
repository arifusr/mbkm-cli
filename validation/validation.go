package validation

import (
	"errors"
	"fmt"

	"github.com/arifusr/mbkm-cli/command"
)

type Validator struct {
	Command *command.Command
}

func NewValidator(c *command.Command) *Validator {
	return &Validator{
		Command: c,
	}
}

func (v *Validator) ValidateCommand(cmd []string) error {
	if len(cmd) < 2 {
		fmt.Println("use the available command")
		fmt.Println("migrate:generate")
		fmt.Println("migrate:run")
		return errors.New("expected args 1")
	}
	for key := range v.Command.CommandAvaliable {
		if key == cmd[1] {
			return nil
		}
	}
	fmt.Println("use the available command")
	fmt.Println("migrate:generate")
	fmt.Println("migrate:run")
	return errors.New("command not found")
}
