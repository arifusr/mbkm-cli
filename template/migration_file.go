package template

const MigrationFile = `
package migration

import (
	"fmt"

	"gorm.io/gorm"
)

type SimpleModel struct {
}

func (s *SimpleModel) TableName() string {
	return "tbl_simple_model"
}

type Aaa struct {
}

func (m *Aaa) Up(db *gorm.DB) error {
	err := db.AutoMigrate(SimpleModel{})
	fmt.Print(err)
	return nil
}
`
