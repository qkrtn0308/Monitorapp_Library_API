package f

import (
	"Library_api/model"
	"Library_api/store/sqlstore"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func Login(c echo.Context) error {
	bodydata, _ := io.ReadAll(c.Request().Body)
	defer c.Request().Body.Close()

	data := model.User{}
	json.Unmarshal(bodydata, &data)
	/****************데이터굴려**************/
	var DB = sqlstore.DBopen()

	row, err := DB.Query("SELECT * FROM userdata where email = $1 AND password = $2", data.Email, data.Password)
	if err != nil {
		panic(err)
	}

	var es []model.User
	var e model.User
	for row.Next() {
		err = row.Scan(&e.ID, &e.FirstName, &e.LastName, &e.Phone, &e.Email, &e.Password, &e.UserStatus, &e.Team)
		if err != nil {
			panic(err)
		}
		es = append(es, e)
	}
	if es == nil {
		a := ("이메일 비밀번호 틀림")
		return c.String(200, a)
	}
	b := e.Team + "팀" + e.LastName + " " + e.FirstName + "님 환영합니다!"

	defer row.Close()
	defer DB.Close()
	return c.String(200, b)
}

func Logout(c echo.Context) error {
	return c.String(http.StatusOK, "로그아웃되었습니다.")
}

func Signup(c echo.Context) error {
	/**************데이터 받음****************/
	bodydata, _ := io.ReadAll(c.Request().Body)
	defer c.Request().Body.Close()

	data := model.User{}
	json.Unmarshal(bodydata, &data)
	/****************데이터굴려**************/
	var db = sqlstore.DBopen()

	insertDynStmt := `insert into "userdata"("firstname", "lastname", "email", "password", "phone_num") values($1, $2, $3, $4, $5)`
	_, er := db.Exec(insertDynStmt, data.FirstName, data.LastName, data.Email, data.Password, data.Phone)
	log.Println(data.ID)
	if er != nil {
		panic(er)
	}
	defer db.Close()

	return c.NoContent(200)
}
