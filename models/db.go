package models

import (
	"fmt"
	"slot/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)



var err error

func ConnectToDb() error{
	dbconfig := config.BuildDBConfig()

	config.DB, err = gorm.Open("mysql", config.DbUrl(dbconfig))
	

	if err != nil {
		fmt.Printf("Status: %v\n", err)
		return err
	}

	config.DB.AutoMigrate(&Event{})
	config.DB.AutoMigrate(&File{})

	fmt.Println("Database connected!")

	return nil
}
