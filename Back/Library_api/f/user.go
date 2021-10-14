package f

import (
	"Library_api/model"
	"Library_api/store"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

//var s = store.New("memory")
var s = store.New("database")

func UserCreate(c echo.Context) error {
	/**************데이터 받음****************/
	bodydata, _ := io.ReadAll(c.Request().Body)
	defer c.Request().Body.Close()
	data := model.User{}
	json.Unmarshal(bodydata, &data)
	/****************데이터굴려**************/
	err := s.UserCreate(&data)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	model.UserInfoByEmail[data.Email] = &data
	return c.JSON(http.StatusOK, model.UserInfoByEmail[data.Email])
}
func FindUserByKeyword(c echo.Context) error {
	/**************데이터 받음****************/
	u := c.Param("u_name")
	/****************데이터굴려**************/
	_, err := s.UserFindByKeyword(u)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusOK)
}
func FindUserByStatusCode(c echo.Context) error {
	/**************데이터 받음****************/
	u := c.Param("u_name")
	n, _ := strconv.Atoi(u)
	/****************데이터굴려**************/
	_, err := s.UserFindByUserStatus(n)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusOK)
}
func UserUpdate(c echo.Context) error {
	/**************데이터 받음****************/
	u := c.Param("u_name")

	bodydata, _ := io.ReadAll(c.Request().Body)
	defer c.Request().Body.Close()
	data := model.User{}
	json.Unmarshal(bodydata, &data)
	/****************데이터굴려**************/
	_, err := s.UserUpdates(u, &data)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	//json 데이터 담는 부분
	return c.JSON(http.StatusOK, model.UserInfoByEmail[u])
}
func DelUserByKeyword(c echo.Context) error {
	/**************데이터 받음****************/
	u1 := c.QueryParam("email")
	u2 := c.QueryParam("password")
	/****************데이터굴려**************/
	err := s.UserDelByKeyword(u1, u2)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.String(http.StatusOK, "삭제")
}
func DelUserByStatusCode(c echo.Context) error {
	err := s.UserDelByStatusCode()
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.String(200, "삭제")
}
