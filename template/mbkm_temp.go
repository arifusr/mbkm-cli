package template

const MbkmTemp = `package main

import (
	"fmt"
	"os"

	migration "github.com/arifusr/mbkm-cli/script/migration/{{.FileName}}"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	// load .env
	if err := godotenv.Load(); err != nil {
		fmt.Print("error load env")
		return
	}
	db, _ := BuildDb()
	model := &migration.NewMigration()
	model.Up(db)
}

func BuildDb() (*gorm.DB, error) {
	sqlCfg := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_USERNAME"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_NAME"),
	)

	db, err := gorm.Open(postgres.Open(sqlCfg), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("failed to connect database")
	}
	return db, err
}`
