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
	return c.NoContent(http.StatusOK)
}
func BookUpdate(c echo.Context) error {
	/**************데이터 받음****************/
	b := c.Param("b_id")

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
	return c.NoContent(http.StatusOK)
}
func FindBookByStatus(c echo.Context) error {
	/**************데이터 받음****************/
	b := c.QueryParam("status")
	log.Printf("%s", b)
	n, _ := strconv.Atoi(b)

	_, err := s.BookFindByBookStatus(n)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.String(http.StatusOK, b)
}
func FindBookByKeyword(c echo.Context) error {
	b := c.Param("b_id")
	if len(b) < 3 {
		return c.String(404, "검색할 키워드가 너무 짧습니다.")
	}
	log.Println(c.Request().URL.String())

	err := s.BookFindByKeyword(b)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
}
func DelBookBykeyword(c echo.Context) error {
	b := c.Param("b_id")

	err := s.BookDelByKeyword(b)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.String(200, b)
}
func DelBookByStatusCode(c echo.Context) error {
	err := s.BookDelByStatusCode()
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.String(200, "삭제")
}
