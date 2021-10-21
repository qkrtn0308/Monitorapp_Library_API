package f

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func Rent(c echo.Context) error {
	L1 := c.QueryParam("userid")
	L2 := c.QueryParam("bookid")
	log.Printf("%s %s", L1, L2)

	err := s.Rent(L1, L2)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusOK)
}
func Return(c echo.Context) error {
	L1 := c.QueryParam("userid")
	L2 := c.QueryParam("bookid")
	log.Printf("%s %s", L1, L2)

	err := s.Return(L1, L2)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusOK)
}
func FindRentByID(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}
func FindReturnByID(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}
