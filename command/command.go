package command

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"

	texttemplate "text/template"

	"github.com/arifusr/mbkm-cli/file"
	"github.com/arifusr/mbkm-cli/model"
	"github.com/arifusr/mbkm-cli/template"
	"github.com/iancoleman/strcase"
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
	commandAvaliable["version"] = command.GetVersion
	commandAvaliable["migrate:undo"] = command.Undo
	command.CommandAvaliable = commandAvaliable
	return command
}

func (c *Command) Undo() error {
	// read last history
	var histories []model.MigrationHistory
	c.DB.Model(&model.MigrationHistory{}).Find(&histories)
	if len(histories) < 1 {
		return nil
	}
	lasthistory := histories[len(histories)-1]
	// get ModuleName
	f, err := os.Open("go.mod")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	b, err := ioutil.ReadAll(f)
	gomod := string(b)

	var regex, _ = regexp.Compile(`module (.*)`)

	var str = regex.FindStringSubmatch(gomod)
	ModuleName := str[1]
	// running Down
	data := struct {
		FileStruct string
		ModuleName string
	}{
		FileStruct: strcase.ToCamel(strings.Replace(lasthistory.MigrationID[20:], ".go", "", 1)),
		ModuleName: ModuleName,
	}
	t, _ := texttemplate.New("mbkm-down").Parse(template.MbkmDown)
	var tpl bytes.Buffer
	t.Execute(&tpl, data)
	// create mbkm_temp.go
	fa := file.NewFile()
	fa.SetDirPath("./")
	fa.SetContent(tpl.String())
	fa.SetName("mbkm_temp.go")

	fa.WriteFile()
	// run exec
	cmd := exec.Command("go", "run", "./mbkm_temp.go")

	out, err := cmd.Output()

	if err != nil {
		fmt.Print(err.Error())

		log.Fatal(err)

	}
	fmt.Print(string(out))

	// jika berahsil migrate hapus history id
	if err := c.DB.Where("migration_id = ?", lasthistory.MigrationID).Delete(&lasthistory).Error; err != nil {
		fmt.Print(err.Error())
	}

	return nil
}

func (c *Command) GetVersion() error {
	fmt.Print("v1.1.2")
	return nil
}

func (c *Command) MigrationRun() error {
	// get migration history
	var histories []model.MigrationHistory
	c.DB.Model(&model.MigrationHistory{}).Find(&histories)
	contain := make(map[string]bool)
	for _, hi := range histories {
		contain[hi.MigrationID] = true
	}
	folder := file.NewFolder()
	files := folder.GetListFile()

	// get ModuleName
	f, err := os.Open("go.mod")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	b, err := ioutil.ReadAll(f)
	gomod := string(b)

	var regex, _ = regexp.Compile(`module (.*)`)

	var str = regex.FindStringSubmatch(gomod)
	ModuleName := str[1]
	// setiap file yg ditemukan running migrate
	for _, fi := range files {
		if contain[fi] {
			continue
		}
		t, _ := texttemplate.New("mbkm-temp").Parse(template.MbkmTemp)
		data := struct {
			FileStruct string
			ModuleName string
		}{
			FileStruct: strcase.ToCamel(strings.Replace(fi[20:], ".go", "", 1)),
			ModuleName: ModuleName,
		}
		var tpl bytes.Buffer
		t.Execute(&tpl, data)
		// create mbkm_temp.go
		f := file.NewFile()
		f.SetDirPath("./")
		f.SetContent(tpl.String())
		f.SetName("mbkm_temp.go")

		f.WriteFile()
		// run exec
		cmd := exec.Command("go", "run", "./mbkm_temp.go")

		out, err := cmd.Output()

		if err != nil {
			fmt.Print(err.Error())

			log.Fatal(err)

		}
		fmt.Print(string(out))

		// jika berahsil migrate tambahkan  file ke migration history
		if err := c.DB.Create(&model.MigrationHistory{
			MigrationID: fi,
		}).Error; err != nil {
			fmt.Print(err.Error())
		}
	}
	// setelah iterasi hapus mbkm_temp
	os.Remove("mbkm_temp.go")

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
	re := regexp.MustCompile("^[a-zA-Z_]*$")
	if !re.MatchString(c.Args[3]) {
		fmt.Println("only underscore and alpha allowed")
		return errors.New("expected options")
	}

	// create file with signature of date
	now := time.Now()
	t, _ := texttemplate.New("mbkm-temp").Parse(template.MigrationFile)
	data := struct {
		FileName string
	}{
		FileName: strcase.ToCamel(c.Args[3]),
	}
	var tpl bytes.Buffer
	t.Execute(&tpl, data)
	filename := now.Format("2006_01_02_15_04_05_") + c.Args[3] + ".go"
	c.File.SetContent(tpl.String())
	c.File.SetName(filename)
	c.File.SetDirPath(os.Getenv("MIGRATION_DIRECTORY") + "/")
	c.File.WriteFile()

	return nil
}
