package sqlstore

import (
	"Library_api/model"
	"log"
)

type SQLstore struct {
}

/*********유저*********/
func (s *SQLstore) UserCreate(data *model.User) error {
	var db = DBopen()

	insertDynStmt := `insert into "userdata"("firstname", "lastname", "email", "password", "phone_num", "status") values($1, $2, $3, $4, $5, $6)`
	_, er := db.Exec(insertDynStmt, data.FirstName, data.LastName, data.Email, data.Password, data.Phone, data.UserStatus)

	if er != nil {
		panic(er)
	}
	defer db.Close()

	return nil
}
func (d *SQLstore) UserFindByUsername(u string) (*model.User, error) {
	var DB = DBopen()

	row, err := DB.Query("SELECT * FROM userdata where username = $1", u)
	if err != nil {
		panic(err)
	}

	var es []model.User
	for row.Next() {
		var e model.User
		err = row.Scan(&e.ID, &e.FirstName, &e.LastName, &e.Email, &e.Password, &e.Phone, &e.UserStatus)
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
func (d *SQLstore) UserUpdates(u string, data *model.User) (*model.User, error) {
	var db = DBopen()

	updateDynStmt := "update userdata set firstname = $1, lastname = $2, email = $3, password = $4, phone_num = $5, status = $6 where email = $7"
	_, er := db.Exec(updateDynStmt, data.FirstName, data.LastName, data.Email, data.Password, data.Phone, data.UserStatus, u)

	if er != nil {
		panic(er)
	}

	defer db.Close()
	return nil, nil
}
func (d *SQLstore) UserDelByName(u string) error {
	var DB = DBopen()

	row, err := DB.Query("DELETE FROM userdata where username = $1", u)
	if err != nil {
		panic(err)
	}

	for row.Next() {
		var e model.User
		err = row.Scan(&e.ID, &e.FirstName, &e.LastName, &e.Email, &e.Password, &e.Phone, &e.UserStatus)
		if err != nil {
			panic(err)
		}
	}
	log.Println("DELETE")

	defer row.Close()
	defer DB.Close()

	return nil

}

/*********책**********/
func (s *SQLstore) BookCreate(data *model.Book) error {
	var db = DBopen()

	insertDynStmt := `insert into "bookdata"("title", "author", "code", "codename, "quantity") values($1, $2, $3, $4, $5)`
	_, er := db.Exec(insertDynStmt, data.Title, data.Author, data.CodeID, data.CodeName, data.Quantity)

	if er != nil {
		panic(er)
	}

	defer db.Close()
	return nil
}
func (s *SQLstore) BookUpdates(b string, data *model.Book) (*model.Book, error) { //todo
	var db = DBopen()
	log.Println(b)
	updateDynStmt := "UPDATE bookdata set title = $1, author = $2, code = $3, codename = $4, status = $5, quantity = $6 where title = $7"
	_, er := db.Exec(updateDynStmt, data.Title, data.Author, data.CodeID, data.CodeName, data.Status, data.Quantity, b)

	if er != nil {
		panic(er)
	}

	defer db.Close()
	return nil, nil
}
func (s *SQLstore) BookFindByBookStatus(b string) (*model.Book, error) {
	log.Println(b)
	var DB = DBopen()
	if b == "" {
		log.Println("not found")
	} else if b != "available" && b != "not_ available" && b != "available,not_available" && b != "not_available,available" {
		log.Println("worng query")
	}
	if b == "available,not_available" || b == "not_available,available" {
		row, err := DB.Query("SELECT * FROM bookdata where status = $1 OR status = $2", "available", "not_available")
		if err != nil {
			panic(err)
		}

		var es []model.Book
		for row.Next() {
			var e model.Book
			err = row.Scan(&e.ID, &e.Title, &e.Author, &e.CodeID, &e.CodeName, &e.Status, &e.Quantity)
			if err != nil {
				panic(err)
			}
			es = append(es, e)
		}
		log.Printf("%v", es)

		defer row.Close()
		defer DB.Close()

		return nil, nil
	}

	row, err := DB.Query("SELECT * FROM bookdata where status = $1 ", b)
	if err != nil {
		panic(err)
	}

	var es []model.Book
	for row.Next() {
		var e model.Book
		err = row.Scan(&e.ID, &e.Title, &e.Author, &e.CodeID, &e.CodeName, &e.Status, &e.Quantity)
		if err != nil {
			panic(err)
		}
		es = append(es, e)
	}
	log.Printf("%v", es)

	defer row.Close()
	defer DB.Close()

	return nil, nil
}
func (s *SQLstore) BookFindByID(n int) error { //todo
	var DB = DBopen()
	log.Println(n)
	row, err := DB.Query("SELECT * FROM bookdata where id = $1", n)
	if err != nil {
		panic(err)
	}

	var es []model.Book
	for row.Next() {
		var e model.Book
		err = row.Scan(&e.ID, &e.Title, &e.Author, &e.CodeID, &e.CodeName, &e.Status, &e.Quantity)
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
func (s *SQLstore) BookDelByID(n int) error { //todo
	var DB = DBopen()
	log.Println(n)
	row, err := DB.Query("DELETE FROM bookdata where id = $1", n)
	if err != nil {
		panic(err)
	}

	for row.Next() {
		var e model.Book
		err = row.Scan(&e.ID, &e.CodeID, &e.CodeName, &e.Title, &e.Author, &e.Status, &e.Quantity)
		if err != nil {
			panic(err)
		}
	}
	log.Println("DELETE")

	defer row.Close()
	defer DB.Close()

	return nil
}

/*********빌림*********/
func (s *SQLstore) OrderCreate(data *model.Order) error {

	var db = DBopen()

	insertDynStmt := `insert into "orderdata" ("bookid", "quantity") values($1, $2)`
	_, er := db.Exec(insertDynStmt, data.BookID, data.Quantity)

	if er != nil {
		panic(er)
	}

	defer db.Close()

	return nil
}
func (s *SQLstore) OrderFindById(n int) error {
	var DB = DBopen()
	log.Println(n)
	row, err := DB.Query("SELECT * FROM orderdata where orderid = $1", n)
	if err != nil {
		panic(err)
	}

	var es []model.Order
	for row.Next() {
		var e model.Order

		err = row.Scan(&e.ID, &e.BookID, &e.Quantity, &e.Status, &e.Complete, &e.RentalDate)

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
func (s *SQLstore) OrderDelById(n int) error {

	var DB = DBopen()

	row, err := DB.Query("DELETE FROM orderdata where orderid = $1", n)
	if err != nil {
		panic(err)
	}

	for row.Next() {
		var e model.Order
		err = row.Scan(&e.ID, &e.BookID, &e.Quantity, &e.Status, &e.Complete, &e.RentalDate)
		if err != nil {
			panic(err)
		}
	}
	log.Println("DELETE")

	defer row.Close()
	defer DB.Close()

	return nil
}
