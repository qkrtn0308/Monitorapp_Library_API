package memstore

import (
	"Library_api/model"
	"errors"
)

type Memstore struct {
}

/*********유저*********/
func (s *Memstore) UserCreate(data *model.User) error {
	if data == nil {
		return errors.New("405")
	} else if model.UserInfoByName[data.Username] != nil {
		return errors.New("405")
	}

	model.Lock.Lock()
	UIDnum++
	data.ID = UIDnum
	data.UserStatus = 1
	model.Lock.Unlock()
	return nil
}
func (s *Memstore) UserFindByUsername(u string) (*model.User, error) {
	if model.UserInfoByName[u] == nil {
		return nil, errors.New("404")
	} else if u == model.UserInfoByName[u].Username {
		return nil, nil
	}
	return nil, errors.New("404")
}
func (s *Memstore) UserUpdates(u string, data *model.User) (*model.User, error) {
	if model.UserInfoByName[u] == nil {
		return nil, errors.New("405")
	}
	//! id , username 고정
	org_id := model.UserInfoByName[u].ID
	org_username := model.UserInfoByName[u].Username
	data.ID = org_id
	data.Username = org_username
	return nil, nil
}
func (s *Memstore) UserDelByName(u string) error {
	if model.UserInfoByName[u] == nil {
		return errors.New("405")
	}
	delete(model.UserInfoByName, u)
	return nil
}

/*********책**********/
func (s *Memstore) BookCreate(data *model.Book) error {
	if data == nil {
		return errors.New("405")
	} else if model.BookInfoByID[data.ID] != nil {
		return errors.New("405")
	} else if model.BookInfoByTitle[data.Title] != nil {
		return errors.New("405")
	}
	model.Lock.Lock()
	BIDnum++
	data.ID = BIDnum
	model.Lock.Unlock()
	return nil
}
func (s *Memstore) BookUpdates(b string, data *model.Book) (*model.Book, error) {
	if model.BookInfoByID[data.ID].ID == data.ID {
		model.BookInfoByID[data.ID] = data
		return nil, nil
	} else if data.ID == 0 {
		return nil, errors.New("405")
	}
	return nil, errors.New("405")
}

func (s *Memstore) BookFindByBookStatus(b string) (*model.Book, error) {
	switch b {
	case "available":
		model.JsonBookArr = nil
		for j := 1; j <= int(len(model.BookInfoByID)); j++ {
			if model.BookInfoByID[int64(j)].Status == "available" {
				model.JsonBookArr = append(model.JsonBookArr, model.BookInfoByID[int64(j)])
			}
		}
		return nil, nil

	case "not_available":
		model.JsonBookArr = nil
		for j := 1; j <= int(len(model.BookInfoByID)); j++ {
			if model.BookInfoByID[int64(j)].Status == "not_available" {
				model.JsonBookArr = append(model.JsonBookArr, model.BookInfoByID[int64(j)])
			}
		}
		return nil, nil

	case "available,not_available", "not_available,available":
		model.JsonBookArr = nil
		for j := 1; j <= int(len(model.BookInfoByID)); j++ {
			model.JsonBookArr = append(model.JsonBookArr, model.BookInfoByID[int64(j)])
		}
		return nil, nil

	default:
		return nil, errors.New("404")
	}
}
func (s *Memstore) BookFindByID(n int) error {
	if n == 0 {
		return errors.New("400")
	} else if model.BookInfoByID[int64(n)] == nil {
		return errors.New("404")
	}

	return nil
}
func (s *Memstore) BookDelByID(n int) error {
	if n == 0 {
		return errors.New("400")
	} else if model.BookInfoByID[int64(n)] == nil {
		return errors.New("404")
	}
	return nil
}

/*********빌림*********/
func (s *Memstore) OrderCreate(data *model.Order) error {
	if data == nil {
		return errors.New("405")
	} else if data.BookID != 0 && data.Quantity != 0 {
		model.Lock.Lock()
		OIDnum++
		data.ID = OIDnum
		data.Status = "placed"
		model.Lock.Unlock()
		return nil
	}

	return errors.New("405")
}
func (s *Memstore) OrderFindById(n int) error {
	if n == 0 {
		return errors.New("400")
	} else if model.OrderInfoByID[int64(n)] == nil {
		return errors.New("404")
	}
	return nil
}
func (s *Memstore) OrderDelById(n int) error {
	if n == 0 {
		return errors.New("400")
	} else if model.OrderInfoByID[int64(n)] == nil {
		return errors.New("404")
	}

	return nil
}
