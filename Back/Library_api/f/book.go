package f

import (
	"Library_api/model"
	"encoding/json"
	"io"
	"log"
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
	b := c.Param("b_name")

	log.Println(c.Request().URL.String())
	log.Println(b)

	bodydata, _ := io.ReadAll(c.Request().Body)
	defer c.Request().Body.Close()
	data := model.Book{}
	json.Unmarshal(bodydata, &data)
	/****************데이터굴려**************/
	_, err := s.BookUpdates(b, &data)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	//json 데이터 담는 부분
	return c.JSON(http.StatusOK, model.UserInfoByName[b])
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
	b := c.Param("b_id")
	log.Println(c.Request().URL.String())
	n, _ := strconv.Atoi(b)

	err := s.BookFindByID(n)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, model.BookInfoByID[int64(n)])
}
func DelBookByID(c echo.Context) error {
	b := c.Param("b_id")

	n, _ := strconv.Atoi(b)

	err := s.BookDelByID(n)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.String(200, b)
}
