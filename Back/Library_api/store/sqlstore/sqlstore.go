package sqlstore

import (
	"Library_api/model"
	"errors"
	"log"
	"time"
)

type SQLstore struct {
}

/*********유저*********/
func (s *SQLstore) UserCreate(data *model.User) error {
	var db = DBopen()

	insertDynStmt := `insert into "userdata"("id", "firstname", "lastname", "team", "phone_num",  "email", "password" ) values($1, $2, $3, $4, $5, $6, $7)`
	_, er := db.Exec(insertDynStmt, data.ID, data.FirstName, data.LastName, data.Team, data.Phone, data.Email, data.Password)

	if er != nil {
		panic(er)
	}
	defer db.Close()

	return nil
}
func (d *SQLstore) UserFindByKeyword(u string) (*model.User, error) {
	var DB = DBopen()

	row, err := DB.Query("SELECT * FROM userdata where id LIKE '%' || $1 || '%' OR team LIKE '%' || $1 || '%' OR firstname LIKE '%' || $1 || '%' OR lastname LIKE '%' || $1 || '%' OR email LIKE '%' || $1 || '%'", u)
	if err != nil {
		panic(err)
	}

	var es []model.User
	for row.Next() {
		var e model.User
		err = row.Scan(&e.ID, &e.FirstName, &e.LastName, &e.Phone, &e.Email, &e.Password, &e.UserStatus, &e.Team)
		if err != nil {
			panic(err)
		}
		es = append(es, e)
	}
	if es == nil {
		log.Println("Can't find anything")
	} else {
		log.Printf("%v", es)
	}

	defer row.Close()
	defer DB.Close()

	return nil, nil
}
func (s *SQLstore) UserFindByUserStatus(u int) (*model.User, error) {
	log.Println(u)
	var DB = DBopen()
	switch u {
	case 0:
		row, err := DB.Query("SELECT * FROM userdata where status = 0")
		if err != nil {
			panic(err)
		}
		var es []model.User
		for row.Next() {
			var e model.User
			err = row.Scan(&e.ID, &e.FirstName, &e.LastName, &e.Phone, &e.Email, &e.Password, &e.UserStatus, &e.Team)
			if err != nil {
				panic(err)
			}
			es = append(es, e)
		}
		if es == nil {
			log.Println("Can't find anything")
		} else {
			log.Printf("%v", es)
		}

		defer row.Close()
		defer DB.Close()

		return nil, nil
	case 1:
		row, err := DB.Query("SELECT * FROM userdata where status = 1")
		if err != nil {
			panic(err)
		}
		var es []model.User
		for row.Next() {
			var e model.User
			err = row.Scan(&e.ID, &e.FirstName, &e.LastName, &e.Phone, &e.Email, &e.Password, &e.UserStatus, &e.Team)
			if err != nil {
				panic(err)
			}
			es = append(es, e)
		}
		if es == nil {
			log.Println("Can't find anything")
		} else {
			log.Printf("%v", es)
		}

		defer row.Close()
		defer DB.Close()

		return nil, nil
	case 2:
		row, err := DB.Query("SELECT * FROM userdata where status = 2")
		if err != nil {
			panic(err)
		}
		var es []model.User
		for row.Next() {
			var e model.User
			err = row.Scan(&e.ID, &e.FirstName, &e.LastName, &e.Phone, &e.Email, &e.Password, &e.UserStatus, &e.Team)
			if err != nil {
				panic(err)
			}
			es = append(es, e)
		}
		if es == nil {
			log.Println("Can't find anything")
		} else {
			log.Printf("%v", es)
		}

		defer row.Close()
		defer DB.Close()

		return nil, nil
	case 3:
		row, err := DB.Query("SELECT * FROM userdata where status = 3")
		if err != nil {
			panic(err)
		}
		var es []model.User
		for row.Next() {
			var e model.User
			err = row.Scan(&e.ID, &e.FirstName, &e.LastName, &e.Phone, &e.Email, &e.Password, &e.UserStatus, &e.Team)
			if err != nil {
				panic(err)
			}
			es = append(es, e)
		}
		if es == nil {
			log.Println("Can't find anything")
		} else {
			log.Printf("%v", es)
		}

		defer row.Close()
		defer DB.Close()

		return nil, nil
	default:
		log.Println("worng query")
		return nil, errors.New("이상하다")
	}
}
func (d *SQLstore) UserUpdates(u string, data *model.User) (*model.User, error) {
	var db = DBopen()

	updateDynStmt := "update userdata set firstname = $1, lastname = $2, email = $3, password = $4, phone_num = $5, status = $6, team = $6   where id = $8 "
	_, er := db.Exec(updateDynStmt, data.FirstName, data.LastName, data.Email, data.Password, data.Phone, data.UserStatus, data.Team, u)

	if er != nil {
		panic(er)
	}

	defer db.Close()
	return nil, nil
}
func (d *SQLstore) UserDelByKeyword(u1 string, u2 string) error {
	var DB = DBopen()

	row, err := DB.Query("DELETE FROM userdata where email = $1 AND password = $2 ", u1, u2)
	if err != nil {
		panic(err)
	}

	for row.Next() {
		var e model.User
		err = row.Scan(&e.ID, &e.FirstName, &e.LastName, &e.Email, &e.Password, &e.Phone, &e.UserStatus, &e.Team)
		if err != nil {
			panic(err)
		}
	}
	log.Println("DELETE")

	defer row.Close()
	defer DB.Close()

	return nil
}
func (s *SQLstore) UserDelByStatusCode() error {
	var DB = DBopen()
	row, err := DB.Query("DELETE FROM userdata where status = 3")
	if err != nil {
		panic(err)
	}

	for row.Next() {
		var e model.User
		err = row.Scan(&e.ID, &e.FirstName, &e.LastName, &e.Phone, &e.Email, &e.Password, &e.UserStatus, &e.Team)
		if err != nil {
			panic(err)
		}
	}
	log.Println("DELETED")

	defer row.Close()
	defer DB.Close()

	return nil
}

/*********책**********/
func (s *SQLstore) BookCreate(data *model.Book) error {
	var db = DBopen()
	insertDynStmt := `insert into "bookdata"("id", "title", "author", "code", "codename", "status","quantity", "releasedate") values($1, $2, $3, $4, $5, $6, $7)`
	_, er := db.Exec(insertDynStmt, data.ID, data.Title, data.Author, data.CodeID, data.CodeName, data.Status, data.Quantity)

	if er != nil {
		panic(er)
	}

	defer db.Close()
	return nil
}
func (s *SQLstore) BookUpdates(b string, data *model.Book) (*model.Book, error) {
	var db = DBopen()
	log.Println(b)
	updateDynStmt := "UPDATE bookdata set id =$1, title = $2, author = $3, code = $4, codename = $5, status = $6, quantity = $7, releasedate = $8 where id = $9"
	_, er := db.Exec(updateDynStmt, data.ID, data.Title, data.Author, data.CodeID, data.CodeName, data.Status, data.Quantity, data.ReleaseDate, b)

	if er != nil {
		panic(er)
	}

	defer db.Close()
	return nil, nil
}
func (s *SQLstore) BookFindByBookStatus(b int, b2 string) ([]model.Book, error) {
	log.Println(b)
	var DB = DBopen()

	switch b {
	case 0:
		row, err := DB.Query("SELECT * FROM bookdata where status = 0 ORDER By $1", b2)
		if err != nil {
			log.Println("b2가 이상하다 ASC/DESC")
			panic(err)
		}
		var es []model.Book
		for row.Next() {
			var e model.Book
			err = row.Scan(&e.ID, &e.Title, &e.Author, &e.CodeID, &e.CodeName, &e.Status, &e.Quantity, &e.ReleaseDate)
			if err != nil {
				panic(err)
			}
			es = append(es, e)
		}
		if es == nil {
			log.Println("Can't find anything")
		} else {
			log.Printf("%v", es)
		}

		defer row.Close()
		defer DB.Close()

		return es, nil
	case 1:
		row, err := DB.Query("SELECT * FROM bookdata where status = 1 ORDER By $1", b2)
		if err != nil {
			log.Println("b2가 이상하다 ASC/DESC")
			panic(err)
		}
		var es []model.Book
		for row.Next() {
			var e model.Book
			err = row.Scan(&e.ID, &e.Title, &e.Author, &e.CodeID, &e.CodeName, &e.Status, &e.Quantity, &e.ReleaseDate)
			if err != nil {
				panic(err)
			}
			es = append(es, e)
		}
		if es == nil {
			log.Println("Can't find anything")
		} else {
			log.Printf("%v", es)
		}

		defer row.Close()
		defer DB.Close()

		return es, nil
	case 2:
		row, err := DB.Query("SELECT * FROM bookdata where status = 2 ORDER By $1", b2)
		if err != nil {
			log.Println("b2가 이상하다 ASC/DESC")
			panic(err)
		}
		var es []model.Book
		for row.Next() {
			var e model.Book
			err = row.Scan(&e.ID, &e.Title, &e.Author, &e.CodeID, &e.CodeName, &e.Status, &e.Quantity, &e.ReleaseDate)
			if err != nil {
				panic(err)
			}
			es = append(es, e)
		}
		if es == nil {
			log.Println("Can't find anything")
		} else {
			log.Printf("%v", es)
		}

		defer row.Close()
		defer DB.Close()

		return es, nil
	case 3:
		row, err := DB.Query("SELECT * FROM bookdata where status = 3 ORDER By $1", b2)
		if err != nil {
			log.Println("b2가 이상하다 ASC/DESC")
			panic(err)
		}
		var es []model.Book
		for row.Next() {
			var e model.Book
			err = row.Scan(&e.ID, &e.Title, &e.Author, &e.CodeID, &e.CodeName, &e.Status, &e.Quantity, &e.ReleaseDate)
			if err != nil {
				panic(err)
			}
			es = append(es, e)
		}
		if es == nil {
			log.Println("Can't find anything")
		} else {
			log.Printf("%v", es)
		}

		defer row.Close()
		defer DB.Close()

		return es, nil
	case 4:
		row, err := DB.Query("SELECT * FROM bookdata where status = 4 ORDER By $1", b2)
		if err != nil {
			log.Println("b2가 이상하다 ASC/DESC")
			panic(err)
		}
		var es []model.Book
		for row.Next() {
			var e model.Book
			err = row.Scan(&e.ID, &e.Title, &e.Author, &e.CodeID, &e.CodeName, &e.Status, &e.Quantity, &e.ReleaseDate)
			if err != nil {
				panic(err)
			}
			es = append(es, e)
		}
		if es == nil {
			log.Println("Can't find anything")
		} else {
			log.Printf("%v", es)
		}

		defer row.Close()
		defer DB.Close()

		return es, nil
	default:
		log.Println("worng query")
		return nil, errors.New("이상하다")
	}
}
func (s *SQLstore) BookFindByKeyword(n string, b2 string) error {
	var DB = DBopen()
	log.Println(b2)

	row, err := DB.Query("SELECT * FROM public.bookdata where id LIKE '%' || $1 || '%' OR title LIKE '%' || $1 || '%' OR author LIKE '%' || $1 || '%' ORDER BY $2", n, b2)
	if err != nil {
		panic(err)
	}

	var es []model.Book
	for row.Next() {
		var e model.Book
		err = row.Scan(&e.ID, &e.Title, &e.Author, &e.CodeID, &e.CodeName, &e.Status, &e.Quantity, &e.ReleaseDate)
		if err != nil {
			panic(err)
		}
		es = append(es, e)
	}
	if es == nil {
		log.Println("Can't find anything")
	} else {
		log.Printf("%v", es)
	}

	defer row.Close()
	defer DB.Close()

	return nil
}
func (s *SQLstore) BookDelById(n string) error {
	var DB = DBopen()
	log.Println(n)
	row, err := DB.Query("DELETE FROM bookdata where id = $1", n)
	if err != nil {
		panic(err)
	}

	for row.Next() {
		var e model.Book
		err = row.Scan(&e.ID, &e.CodeID, &e.CodeName, &e.Title, &e.Author, &e.Status, &e.Quantity, &e.ReleaseDate)
		if err != nil {
			panic(err)
		}
	}
	log.Println("DELETED")

	defer row.Close()
	defer DB.Close()

	return nil
}
func (s *SQLstore) BookDelByStatusCode() error {
	var DB = DBopen()
	row, err := DB.Query("DELETE FROM bookdata where status = 4")
	if err != nil {
		panic(err)
	}

	for row.Next() {
		var e model.Book
		err = row.Scan(&e.ID, &e.CodeID, &e.CodeName, &e.Title, &e.Author, &e.Status, &e.Quantity, &e.ReleaseDate)
		if err != nil {
			panic(err)
		}
	}
	log.Println("DELETED")

	defer row.Close()
	defer DB.Close()

	return nil
}

/*********빌림*********/
func (s *SQLstore) Rent(L1 string, L2 string) error {
	var DB = DBopen()
	//! user id 판별
	row, err := DB.Query("SELECT * FROM userdata where id = $1", L1)
	if err != nil {
		panic(err)
	}
	var es []model.User
	for row.Next() {
		var e model.User
		err = row.Scan(&e.ID, &e.FirstName, &e.LastName, &e.Phone, &e.Email, &e.Password, &e.UserStatus, &e.Team)
		if err != nil {
			panic(err)
		}
		es = append(es, e)
	}
	if es == nil {
		log.Println("존재하지 않는 USER 입니다.")
		return err
	}
	//! book id 판별
	row, err = DB.Query("SELECT * FROM bookdata where id = $1", L2)
	if err != nil {
		panic(err)
	}
	var ea []model.Book
	for row.Next() {
		var e model.Book
		err = row.Scan(&e.ID, &e.Title, &e.Author, &e.CodeID, &e.CodeName, &e.Status, &e.Quantity)
		if err != nil {
			panic(err)
		}
		ea = append(ea, e)
	}
	if ea == nil {
		log.Println("존재하지 않는 BOOK 입니다.")
		return err
	}
	//! 데이터 삽입
	stat := 2
	temp := "no return yet"
	t := time.Now().Format("2006-01-02 15:04:05")
	insertDynStmt := `insert into "log" ("userid", "bookid", "rentaldate", "returndate") values($1, $2, $3, $4)`
	_, er := DB.Exec(insertDynStmt, L1, L2, t, temp)
	if er != nil {
		panic(er)
	}
	updateDynStmt := ("UPDATE bookdata set status = $1 where id = $2")
	_, er = DB.Exec(updateDynStmt, stat, L2)
	if er != nil {
		panic(er)
	}

	defer DB.Close()
	return nil

}
func (s *SQLstore) Return(L1 string, L2 string) error {
	temp := "no return yet"
	var DB = DBopen()
	//! id 판별
	row, err := DB.Query("SELECT * FROM log where userid = $1 AND bookid = $2 AND returndate = $3", L1, L2, temp)
	if err != nil {
		panic(err)
	}
	var el []model.Log
	for row.Next() {
		var e model.Log
		err := row.Scan(&e.UserID, &e.BookID, &e.RentalDate, &e.Returndate)
		if err != nil {
			panic(err)
		}
		el = append(el, e)
	}
	log.Printf("%v", el)
	if el == nil {
		log.Println("그런거 업다.")
		return err
	}

	//! 데이터 삽입
	stat := 1
	t := time.Now().Format("2006-01-02 15:04:05")
	updateDynStmt := ("UPDATE log set returndate = $1 where userid = $2 AND bookid = $3")
	_, er := DB.Exec(updateDynStmt, t, L1, L2)
	if er != nil {
		panic(er)
	}

	updateDynStmt = ("UPDATE bookdata set status = $1 where id = $2")
	_, er = DB.Exec(updateDynStmt, stat, L2)

	if er != nil {
		panic(er)
	}

	defer DB.Close()
	return nil
}
func (s *SQLstore) FindRentByID() error {
	return nil
}
func (s *SQLstore) FindReturnByID() error {
	return nil
}
