package main

import (
	"fmt"
	"task3/controllers"
	"task3/services"
	"task3/models"
)

func main(){
	library := &services.Library{
		Members: make(map[int]*models.Member),
		Books:   make(map[int]*models.Book),
	}
	library.Members[1] = &models.Member1

	for {

		fmt.Printf(`
		What service would you like to get.
		(The number associated with the service you want)
		("Encoded member id: 1 name: "Green")
		1. Add a Book
		2. Remove a Book
		3. Borrow a Book
		4. Return a Book
		5. List Available Books
		6. List Borrowed Books
		`)

		service_no := 0
		_, err := fmt.Scanln(&service_no)
		if err != nil{
			fmt.Println("Wrong Input Plese Try Again!")
		} 

		switch service_no {
		case 1:
			controllers.HandleAddBook(library)
		case 2:
			controllers.HandleRemoveBook(library)
		case 3:
			controllers.HandleBorrowBook(library)
		case 4:
			controllers.HandleReturnBook(library)
		case 5:
			controllers.HandleListAvailableBooks(library)
		case 6:
			controllers.HandleListBorrowedBooks(library)
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}