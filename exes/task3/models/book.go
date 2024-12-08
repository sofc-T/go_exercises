package models

import (
	"errors"
)

type Status string


const(
	
	Borrowed string = "Borrowed"
	Available string = "Available"
)

type Book struct{
	Id int 
	Title string
	Author string 
	Status string
}

func (b *Book) SetStatus(s string) error {
	if s != Borrowed && s != Available{
		return errors.New("Invalid status value")
	}
	b.Status = s
	return nil
}

