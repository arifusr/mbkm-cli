package command

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"time"

	"github.com/arifusr/mbkm-cli/file"
	"github.com/arifusr/mbkm-cli/model"
	"gorm.io/gorm"
)

type Command struct {
	CommandAvaliable map[string]func() error
	Args             []string
	File             *file.File
	DB               *gorm.DB
}

func NewCommand(args []string, db *gorm.DB) *Command {
	command := &Command{
		Args: args,
		File: file.NewFile(),
		DB:   db,
	}
	commandAvaliable := make(map[string]func() error)
	commandAvaliable["migrate:generate"] = command.MigrationGenerate
	commandAvaliable["migrate:run"] = command.MigrationRun
	command.CommandAvaliable = commandAvaliable
	return command
}

func (c *Command) MigrationRun() error {
	// get migration history
	var histories []model.MigrationHistory
	c.DB.Model(&model.MigrationHistory{}).Find(&histories)

	folder := file.NewFolder()
	files := folder.GetListFile()
	fmt.Print(files)
	return nil
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
	now := time.Now()
	filename := now.Format("2006_01_02_15_04_05_") + c.Args[3] + ".go"
	c.File.SetContent("aaaa")
	c.File.SetName(filename)
	c.File.SetDirPath(os.Getenv("MIGRATION_DIRECTORY") + "/")
	c.File.WriteFile()

	return nil
}
