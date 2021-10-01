package model

import (
	"sync"
	"time"
)

type Book struct {
	ID     int64    `json:"id"`
	Genre  Category `json:"genre" validate:"required"`
	Title  string   `json:"title" validate:"required, containsrune"`
	Author string   `json:"author" validate:"required, containsrune"`
	Status string   `json:"status" validate:"required, containsrune"`
}
type Category struct {
	ID   int64  `json:"id" validate:"required, numeric"`
	Name string `json:"name" validate:"required, containsrune"`
}
type User struct {
	ID         int64  `json:"id"`
	Username   string `json:"username" validate:"required, containsrune"`
	FirstName  string `json:"firstName" validate:"required, containsrune"`
	LastName   string `json:"lastName" validate:"required, containsrune"`
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"required, containsrune"`
	Phone      string `json:"phone" validate:"required, numeric"`
	UserStatus int32  `json:"userStatus"`
}
type Order struct {
	ID         int64     `json:"id"`
	BookID     int64     `json:"bookId" validate:"required, numeric"`
	Quantity   int32     `json:"quantity" validate:"required, numeric"`
	RentalDate time.Time `json:"rentalDate"`
	Status     string    `json:"status"`
	Complete   bool      `json:"complete"`
}

var BookInfoByID = make(map[int64]*Book)
var BookInfoByTitle = make(map[string]*Book)
var UserInfoByName = make(map[string]*User)
var OrderInfoByID = make(map[int64]*Order)

var JsonBookArr []*Book
var JsonBookStr *Book
var JsonUserStr *User
var JsonOrderStr *Order

var (
	Lock *sync.Mutex = &sync.Mutex{}
)
