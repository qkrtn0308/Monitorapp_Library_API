package store

import (
	"Library_api/model"
	"Library_api/store/sqlstore"
)

type Store interface {
	/*********유저*********/
	UserCreate(*model.User) error
	UserFindByKeyword(string) (*model.User, error)
	UserFindByUserStatus(int) (*model.User, error)
	UserUpdates(string, *model.User) (*model.User, error)
	UserDelByKeyword(string, string) error
	UserDelByStatusCode() error
	/**********책*********/
	BookCreate(*model.Book) error
	BookUpdates(string, *model.Book) (*model.Book, error)
	BookFindByBookStatus(int) (*model.Book, error)
	BookFindByKeyword(string) error
	BookDelByKeyword(string) error
	BookDelByStatusCode() error
	/*********오더*********/
	Rent(string, string) error
	Return(string, string) error
	// OrderFindById(int) error
	// OrderDelById(int) error
}

func New(storeType string) Store {
	switch storeType {
	// case "memory":
	// return &memstore.Memstore{}
	case "database":
		return &sqlstore.SQLstore{}
	default:
		return &sqlstore.SQLstore{}
	}
}
