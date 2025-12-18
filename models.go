package main

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name        string `Json:"name"`
	Author      string `Json:"author"`
	Description string `Json:"description"`
	Price       uint   `Json:"price"`
}

func createBook(db *gorm.DB, book *Book) error {
	result := db.Create(book)
	if result.Error != nil {
		return result.Error
	}
	fmt.Println("Create Books Suscressful ")
	return nil

}
func getBook(db *gorm.DB, id int) *Book {
	var book Book
	result := db.First(&book, id)
	if result.Error != nil {
		log.Fatalf("Error create book: %v", result.Error)
	}
	return &book

}
func getBooks(db *gorm.DB) []Book {
	var books []Book
	result := db.Find(&books)
	if result.Error != nil {
		log.Fatalf("Error create book: %v", result.Error)
	}
	return books

}
func updateBook(db *gorm.DB, book *Book) error {
	//result := db.Model(&book).Update(book)
	result := db.Save(&book)
	if result.Error != nil {
		return result.Error
	}
	fmt.Println("update Successful")
	return nil

}
func deleteBook(db *gorm.DB, id int) error {
	var book Book
	result := db.Delete(&book, id)
	if result.Error != nil {
		return result.Error
	}
	fmt.Println("DeleteBook Successful")
	return nil

}
func searcBook(db *gorm.DB, bookName string) []Book {
	var books []Book
	//first find
	result := db.Where("name = ?", bookName).Order("price").Find(&books)
	if result.Error != nil {
		log.Fatalf("Error Search book: %v", result.Error)
	}
	return books

}
