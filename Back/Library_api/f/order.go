package f

import (
	"Monitorapp_Library/model"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func OrderCreate(c echo.Context) error {
	/**************데이터 받음****************/
	bodydata, _ := io.ReadAll(c.Request().Body)
	defer c.Request().Body.Close()
	data := model.Order{}
	json.Unmarshal(bodydata, &data)
	/****************데이터굴려**************/
	err := s.OrderCreate(&data)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	model.OrderInfoByID[data.ID] = &data
	return c.JSON(http.StatusCreated, model.OrderInfoByID[data.ID])
}
func FindOrderById(c echo.Context) error {
	o := c.Param("o_id")
	n, _ := strconv.Atoi(o)

	err := s.OrderFindById(n)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, model.OrderInfoByID[int64(n)])
}
func DelOrderById(c echo.Context) error {
	o := c.Param("o_id")
	n, _ := strconv.Atoi(o)

	err := s.OrderDelById(n)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	delete(model.OrderInfoByID, int64(n))
	return c.NoContent(http.StatusOK)
}
