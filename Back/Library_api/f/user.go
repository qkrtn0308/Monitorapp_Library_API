package f

import (
	"Monitorapp_Library/model"
	"Monitorapp_Library/store"
	"encoding/json"
	"io"
	"net/http"

	"github.com/labstack/echo"
)

var s = store.New("memory")

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
	model.UserInfoByName[data.Username] = &data
	return c.JSON(http.StatusOK, model.UserInfoByName[data.Username])
}
func FindUserByName(c echo.Context) error {
	/**************데이터 받음****************/
	u := c.Param("u_name")
	/****************데이터굴려**************/
	_, err := s.UserFindByUsername(u)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, model.UserInfoByName[u])
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
	model.UserInfoByName[u] = &data
	return c.JSON(http.StatusOK, model.UserInfoByName[u])
}
func DelUserByName(c echo.Context) error {
	/**************데이터 받음****************/
	u := c.Param("u_name")
	/****************데이터굴려**************/
	err := s.UserDelByName(u)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.String(http.StatusOK, "삭제")
}
