package main

import (
	"fmt"
	"os"

	"github.com/arifusr/mbkm-cli/command"
	"github.com/arifusr/mbkm-cli/config"
	"github.com/arifusr/mbkm-cli/validation"
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
	// load .mbkm-cli
	if err := godotenv.Load(".mbkm-cli"); err != nil {
		fmt.Print("please create .mbkm-cli to setup your migration directory")
		return
	}
	configenv := config.NewConfig()
	if err := configenv.ValidateConfig(); err != nil {
		return
	}
	db, _ := buildDb()
	args := os.Args
	cmd := command.NewCommand(args, db)

	validator := validation.NewValidator(cmd)
	if err := validator.ValidateCommand(args); err != nil {
		return
	}
	//run command
	cmd.CommandAvaliable[args[1]]()
}

func buildDb() (*gorm.DB, error) {
	sqlCfg := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_USERNAME"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_NAME"),
	)

	db, err := gorm.Open(postgres.Open(sqlCfg), &gorm.Config{
		Logger: logger.Default,
	})
	if err != nil {
		panic("failed to connect database")
	}
	return db, err
}
