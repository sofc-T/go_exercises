package models



type LibraryManager interface{
	AddBooK(b *Book) error
	RemoveBook(bid *int) error 
	BorrowBook(bid *int, mid *int) error 
	ReturnBook(bid *int, mid *int) error 
	ListAvailableBooks() []Book
	ListBorrowedBooks(mid *int) []Book
	
}



