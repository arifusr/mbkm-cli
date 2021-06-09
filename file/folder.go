package file

import (
	"io/ioutil"
	"log"
	"os"
)

type Folder struct {
	DirPath string
}

func NewFolder() *Folder {
	return &Folder{
		DirPath: os.Getenv("MIGRATION_DIRECTORY"),
	}
}

func (f *Folder) GetListFile() (res []string) {
	files, err := ioutil.ReadDir("./" + f.DirPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		res = append(res, f.Name())
	}
	return
}
