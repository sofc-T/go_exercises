package services

import (
	"errors"
	"task3/models"
)



type Library struct{
	Members map[int]*models.Member
	Books map[int]*models.Book

}


func (l *Library) AddBook(title string, id int, author string) error {

	if _, exists := l.Books[id]; exists {  
        return errors.New("book already exists")  
    }

	book := models.Book{Id: id, Title: title, Author: author, Status: models.Available}
	l.Books[id] = &book 
	return nil
}

func (l *Library) RemoveBook(id int) error {

	if _, exists := l.Books[id]; !exists {  
        return errors.New("book does not exist") 
    }

	delete(l.Books, id)
	return nil

}

func (l *Library) BorrowBook(bid int, mid int) error{
	member, err := l.Members[mid]
	if !err {
		return errors.New("member doesn't Exist")
	} 

	book, err := l.Books[bid]
	if !err {
		return errors.New("book doesn't Exist")
	}

	if book.Status != models.Available {  
        return errors.New("book is not Availale")
    }

	book.SetStatus(models.Borrowed)
	member.BorrowedBooks[bid] = book
	return nil
	
}

func (l *Library) ReturnBook(bid int, mid int) error {
	member, err := l.Members[mid]
	if !err{
		return errors.New("member doesn't Exist")
	} 

	book, err := l.Books[bid]
	if !err{
		return errors.New("book doesn't Exist")
	}

	if book.Status != models.Borrowed {  
        return errors.New("book was not borrowed")
    }

	book.SetStatus(models.Available)
	delete(member.BorrowedBooks, bid)
	return nil

}


func (l *Library) AddMember(name string , mid int) error {
	_, err := l.Members[mid]
	if !err{
		return errors.New("member already exists")
	}

	member := models.Member{Name: name, Id : mid, BorrowedBooks: make(map[int]*models.Book)}
	l.Members[mid] = &member 
	return nil
}


func (l *Library) RemoveMember(name string, mid int) error {
	_, err := l.Members[mid]
	if !err {
		return errors.New("member doesn't Exist")
	}

	
	delete(l.Members, mid)
	return nil
}

func (l *Library) ListAvailableBooks() []models.Book {
	var availableBooks []models.Book
	for _, book := range l.Books {
		if book.Status == models.Available {
			availableBooks = append(availableBooks, *book)
		}
	}
	return availableBooks
}

func (l *Library) ListBorrowedBooks(memberID int) ([]models.Book, error) {
	member, exists := l.Members[memberID]
	if !exists {
		return nil, errors.New("member does not exist")
	}
	var borrowedBooks []models.Book
	for _, book := range member.BorrowedBooks {
		borrowedBooks = append(borrowedBooks, *book)
	}
	return borrowedBooks, nil
}