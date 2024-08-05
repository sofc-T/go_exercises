package controllers

import (
	"fmt"
	"log"
	"task3/services"
)

func HandleAddBook(library *services.Library) {
	var id int
	var title, author string
	fmt.Print("Enter book ID: ")
	fmt.Scanln(&id)
	fmt.Print("Enter book title: ")
	fmt.Scanln(&title)
	fmt.Print("Enter book author: ")
	fmt.Scanln(&author)

	err := library.AddBook( title, id,  author)
	if err != nil || id < 1{
		log.Println("Error:", "Invalid Input")
	} else {
		fmt.Println("Book added successfully!")
	}
}

func HandleRemoveBook(library *services.Library) {
	var id int
	fmt.Print("Enter book ID to remove: ")
	fmt.Scanln(&id)
	err := library.RemoveBook(id)
	if err != nil {
		log.Println("Error:", err)
	} else {
		fmt.Println("Book removed successfully!")
	}
}

func HandleBorrowBook(library *services.Library) {
	var bookID, memberID int
	fmt.Print("Enter book ID to borrow: ")
	fmt.Scanln(&bookID)
	fmt.Print("Enter member ID: ")
	fmt.Scanln(&memberID)
	err := library.BorrowBook(bookID, memberID)
	if err != nil {
		log.Println("Error:", err)
	} else {
		fmt.Println("Book borrowed successfully!")
	}
}

func HandleReturnBook(library *services.Library) {
	var bookID, memberID int
	fmt.Print("Enter book ID to return: ")
	fmt.Scanln(&bookID)
	fmt.Print("Enter member ID: ")
	fmt.Scanln(&memberID)
	err := library.ReturnBook(bookID, memberID)
	if err != nil {
		log.Println("Error:", err)
	} else {
		fmt.Println("Book returned successfully!")
	}
}

func HandleListAvailableBooks(library *services.Library) {
	books := library.ListAvailableBooks()
	fmt.Println("Available Books:")
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.Id, book.Title, book.Author)
	}
}

func HandleListBorrowedBooks(library *services.Library) {
	var memberID int
	fmt.Print("Enter member ID: ")
	fmt.Scanln(&memberID)
	books, err := library.ListBorrowedBooks(memberID)
	if err != nil {
		log.Println("Error:", err)
		return
	}
	fmt.Println("Borrowed Books:")
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.Id, book.Title, book.Author)
	}
}
