package file

import "io/ioutil"

type File struct {
	Content string
	Name    string
	DirPath string
}

func NewFile() *File {
	return &File{}
}

func (f *File) SetContent(content string) {
	f.Content = content
}

func (f *File) SetName(name string) {
	f.Name = name
}

func (f *File) SetDirPath(path string) {
	f.DirPath = path
}

func (f *File) WriteFile() error {
	err := ioutil.WriteFile(f.DirPath+f.Name, []byte(f.Content), 0644)
	return err
}
