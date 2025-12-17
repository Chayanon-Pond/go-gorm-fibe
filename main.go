package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	host         = "127.0.0.1"
	port         = 5555
	databaseName = "mydatabase"
	username     = "myuser"
	password     = "mypassword"
)

func main() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, username, password, databaseName)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Enable color
		},
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		log.Panic("failed to connect database !")
	}
	fmt.Println("connect Successful !")
	db.AutoMigrate(&Book{})
	currentBook := getBook(db, 1)
	currentBook.Name = "New Pond"
	currentBook.Price = 999
	updateBook(db, currentBook)

}
