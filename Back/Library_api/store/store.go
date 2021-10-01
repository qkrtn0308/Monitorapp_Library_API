package store

import (
	"Library_api/model"
	"Library_api/store/memstore"
)

type Store interface {
	/*********유저*********/
	UserCreate(*model.User) error
	UserFindByUsername(string) (*model.User, error)
	UserUpdates(string, *model.User) (*model.User, error)
	UserDelByName(string) error
	/**********책*********/
	BookCreate(*model.Book) error
	BookUpdates(*model.Book) (*model.Book, error)
	BookFindByBookStatus(string) (*model.Book, error)
	BookFindByID(int) error
	BookDelByID(int) error
	/*********오더*********/
	OrderCreate(*model.Order) error
	OrderFindById(int) error
	OrderDelById(int) error
}

func New(storeType string) Store {
	switch storeType {
	case "memory":
		return &memstore.Memstore{}
	// case "database":
	// return &sqlstore.SQLStore{}
	default:
		return &memstore.Memstore{}
	}
}
