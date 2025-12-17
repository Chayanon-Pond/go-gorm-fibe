package main

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name        string
	Author      string
	Description string
	Price       uint
}

func createBook(db *gorm.DB, book *Book) {
	result := db.Create(book)
	if result.Error != nil {
		log.Fatalf("Error create book: %v", result.Error)
	}
	fmt.Println("Create Books Suscressful ")

}
func getBook(db *gorm.DB, id uint) *Book {
	var book Book
	result := db.First(&book, id)
	if result.Error != nil {
		log.Fatalf("Error create book: %v", result.Error)
	}
	return &book

}
func updateBook(db *gorm.DB, book *Book) {
	result := db.Save(&book)
	if result.Error != nil {
		log.Fatalf("Error create book: %v", result.Error)
	}
	fmt.Println("update Successful")

}
func deleteBook(db *gorm.DB, id uint) {
	var book Book
	result := db.Delete(&book, id)
	if result.Error != nil {
		log.Fatalf("Error create book: %v", result.Error)
	}
	fmt.Println("DeleteBook Successful")

}
