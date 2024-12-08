package models

import (
)


type Member struct{
	Id int 
	Name string
	BorrowedBooks map[int]*Book
}

var Member1 = Member{Id :1 , Name: "Green", BorrowedBooks: make(map[int]*Book)}
