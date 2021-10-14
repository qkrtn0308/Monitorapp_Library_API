package main

import (
	"Library_api/f"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/", f.Signup)
	e.GET("/login", f.Login)
	e.GET("/logout", f.Logout)

	e.POST("/book", f.BookCreate)
	e.GET("/book", f.FindUserByStatusCode)
	e.GET("/book/:b_id", f.FindBookByKeyword)
	e.GET("/book/code", f.FindBookByStatus)
	e.PUT("/book/:b_id", f.BookUpdate)
	e.DELETE("/book/:code4:", f.DelBookByStatusCode)

	e.POST("/user", f.UserCreate)
	e.GET("/user/:u_name", f.FindUserByKeyword)
	e.GET("/user/code", f.FindUserByStatusCode)
	e.PUT("/user/:u_name", f.UserUpdate)
	e.DELETE("/user/key", f.DelUserByKeyword)
	e.DELETE("/user/code3", f.DelUserByStatusCode)

	e.POST("/library/rent", f.Rent)
	e.GET("/library/order/:o_id", f.Return)

	//start server
	e.Start(":4000")

}
