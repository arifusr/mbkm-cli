package template

const MigrationFile = `
package migration

import (
	"gorm.io/gorm"
)

type {{.FileName}} struct {
}

func (m *{{.FileName}}) Up(db *gorm.DB) error {
	// TODO - migrate here
	return nil
}
`
