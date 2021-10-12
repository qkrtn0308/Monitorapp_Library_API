package f

import (
	"Library_api/model"
	"Library_api/store/sqlstore"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

func Login(c echo.Context) error {
	a := c.QueryParam("username")
	b := c.QueryParam("password")

	validator.New()
	if model.UserInfoByName[a] == nil {
		return c.NoContent(http.StatusBadRequest)
	}

	if model.UserInfoByName[a].Username == a && model.UserInfoByName[a].Password == b {
		return c.String(http.StatusOK, "로그인 되었습니다~~")
	} else {
		return c.NoContent(http.StatusBadRequest)
	}
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

	insertDynStmt := `insert into "userdata"("username", "firstname", "lastname", "email", "password", "phone_num") values($1, $2, $3, $4, $5, $6)`
	_, er := db.Exec(insertDynStmt, data.Username, data.FirstName, data.LastName, data.Email, data.Password, data.Phone)
	log.Println(data.ID)
	if er != nil {
		panic(er)
	}
	defer db.Close()

	return c.NoContent(200)
}
