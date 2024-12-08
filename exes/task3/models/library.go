package models


import (

)

type LibraryManager interface{
	AddBooK(b *Book) error
	RemoveBook(bid *int) error 
	BorrowBook(bid *int) error 
	ReturnBook(bid *int) error 
	ListAvailableBooks() error 
	ListBorrowedBooks() error 
	
}



