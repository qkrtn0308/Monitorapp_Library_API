package model

import (
	"sync"
)

type User struct {
	ID         string `json:"id"`
	FirstName  string `json:"firstname" validate:"required, containsrune"`
	LastName   string `json:"lastname" validate:"required, containsrune"`
	Team       string `json:"team" validate:"required"`
	Phone      string `json:"phone" validate:"required"`
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"required, containsrune"`
	UserStatus int32  `json:"userstatus"`
	// 0 : 관리자, 1 : 사원, 2 : 대출금지, 3 : 퇴사(삭제예정) => default : 1
}
type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title" validate:"required, containsrune"`
	Author string `json:"author" validate:"required, containsrune"`
	CodeID string `json:"code" validate:"required"`
	//   000   100    200        300     400   500    600        700     800
	CodeName string `json:"codename" validate:"required, containsrune"`
	// 사내문건 잡지 경제/경영 인문사회과학 과학   IT 연구소전문서적 자기개발 소설
	Status int64 `json:"status" validate:"required, containsrune"`
	// 0 : 추가예정, 1 : 대출가능, 2 : 대출중, 3 : 대출금지(훼손/분실_복구예정), 4 : 대출금지(훼손/분실_삭제예정)
	Quantity int64 `json:"quantity"`
	// default : 1
}
type Log struct {
	UserID     string `json:"userid"`
	BookID     string `json:"bookId" validate:"required, numeric"`
	RentalDate string `json:"rentaldate"`
	Returndate string `json:"returndate"`
	// YYYY-MM-DD hh:mm:ss
	// t := time.Now().Format("YYYY-MM-DD hh:mm:ss")

}

var BookInfoByID = make(map[string]*Book)
var BookInfoByTitle = make(map[string]*Book)
var UserInfoByEmail = make(map[string]*User)
var LogInfoByID = make(map[int64]*Log)

var JsonBookArr []*Book
var JsonBookStr *Book
var JsonUserStr *User
var JsonLogStr *Log

var (
	Lock *sync.Mutex = &sync.Mutex{}
)
