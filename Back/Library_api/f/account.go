package f

import (
	"Library_api/model"
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
