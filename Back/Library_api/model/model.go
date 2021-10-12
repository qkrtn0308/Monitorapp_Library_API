package model

import (
	"sync"
)

type Book struct {
	ID       int64  `json:"id"`
	CodeID   int64  `json:"code" validate:"required, numeric"`
	CodeName string `json:"codename" validate:"required, containsrune"`
	Title    string `json:"title" validate:"required, containsrune"`
	Author   string `json:"author" validate:"required, containsrune"`
	Status   string `json:"status" validate:"required, containsrune"`
}

type User struct {
	ID         int64  `json:"id"`
	Username   string `json:"username" validate:"required, containsrune"`
	FirstName  string `json:"firstname" validate:"required, containsrune"`
	LastName   string `json:"lastname" validate:"required, containsrune"`
	Phone      string `json:"phone" validate:"required"`
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"required, containsrune"`
	UserStatus int32  `json:"userstatus"`
}
type Order struct {
	ID         int64  `json:"id"`
	BookID     int64  `json:"bookId" validate:"required, numeric"`
	Quantity   int32  `json:"quantity" validate:"required, numeric"`
	RentalDate string `json:"rentaldate"`
	Status     string `json:"status"`
	Complete   bool   `json:"complete"`
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
