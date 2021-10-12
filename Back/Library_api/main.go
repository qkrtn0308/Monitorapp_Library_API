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
	e.GET("/book", f.FindBookStatus)
	e.GET("/book/:b_id", f.FindBookByID)
	e.PUT("/book/:b_id", f.BookUpdate)
	e.DELETE("/book/:b_id:", f.DelBookByID)

	e.POST("/user", f.UserCreate)
	e.GET("/user/:u_name", f.FindUserByName)
	e.PUT("/user/:u_name", f.UserUpdate)
	e.DELETE("/user/:u_name", f.DelUserByName)

	e.POST("/library/order", f.OrderCreate)
	e.GET("/library/order/:o_id", f.FindOrderById)
	e.DELETE("/library/order/:o_id", f.DelOrderById)

	//start server
	e.Start(":4000")

}
