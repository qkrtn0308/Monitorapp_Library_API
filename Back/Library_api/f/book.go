package f

import (
	"Monitorapp_Library/model"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func BookCreate(c echo.Context) error {
	/**************데이터 받음****************/
	bodydata, _ := io.ReadAll(c.Request().Body)
	defer c.Request().Body.Close()
	data := model.Book{}
	json.Unmarshal(bodydata, &data)
	/****************데이터굴려**************/
	err := s.BookCreate(&data)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	model.BookInfoByID[data.ID] = &data
	model.BookInfoByTitle[data.Title] = &data
	return c.JSON(http.StatusOK, model.BookInfoByID[data.ID])
}
func BookUpdate(c echo.Context) error {
	/**************데이터 받음****************/
	bodydata, _ := io.ReadAll(c.Request().Body)
	defer c.Request().Body.Close()
	data := model.Book{}
	json.Unmarshal(bodydata, &data)
	/****************데이터굴려**************/
	_, err := s.BookUpdates(&data)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	model.BookInfoByID[data.ID] = &data
	return c.JSON(http.StatusOK, model.BookInfoByID[data.ID])
}
func FindBookStatus(c echo.Context) error {
	/**************데이터 받음****************/
	b := c.QueryParam("status")

	_, err := s.BookFindByBookStatus(b)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, model.JsonBookArr)
}
func FindBookByID(c echo.Context) error {
	p := c.Param("b_id")

	n, _ := strconv.Atoi(p)

	err := s.BookFindByID(n)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, model.BookInfoByID[int64(n)])
}

// func FindBookByTitle(c echo.Context) error {}
func DelBookByID(c echo.Context) error {
	model.JsonBookArr = nil
	p := c.Param("b_id")

	n, _ := strconv.Atoi(p)

	err := s.BookFindByID(n)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	delete(model.BookInfoByID, int64(n))

	return c.NoContent(http.StatusOK)
}
