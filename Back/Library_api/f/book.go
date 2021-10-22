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
	b2 := c.QueryParam("sort")
	n, _ := strconv.Atoi(b)
	log.Printf("%s", b)

	log.Print(b2, n)

	u, err := s.BookFindByBookStatus(n, b2)
	if err != nil {
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		c.Response().Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT")
		return c.NoContent(http.StatusInternalServerError)
	}

	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	c.Response().Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT")
	return c.JSON(http.StatusOK, u)
}
func FindBookByKeyword(c echo.Context) error {
	b := c.QueryParam("keyword")
	b2 := c.QueryParam("sort")
	log.Printf("%s, %s", b, b2)
	if len(b) < 2 {
		return c.String(404, "검색할 키워드가 너무 짧습니다.")
	}
	log.Println(c.Request().URL.String())

	err := s.BookFindByKeyword(b, b2)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
}
func DelBookBykeyword(c echo.Context) error {
	b := c.Param("b_id")

	err := s.BookDelById(b)
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
