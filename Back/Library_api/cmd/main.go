package main

import (
	"Library_api/f"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	//e.Static("/public/", "./public/")
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/a", f.Login)

	e.GET("/user/login", f.Login)
	e.GET("/user/logout", f.Logout)

	e.POST("/book", f.BookCreate)
	e.PUT("/book", f.BookUpdate)
	e.GET("/book/findByStatus", f.FindBookStatus)
	e.GET("/book/:b_id", f.FindBookByID)
	e.DELETE("/book/:b_id:", f.DelBookByID)

	e.POST("/user", f.UserCreate)
	e.GET("/user/:u_name", f.FindUserByName)
	e.PUT("/user/:u_name", f.UserUpdate)
	e.DELETE("/user/:u_name", f.DelUserByName)

	e.POST("/library/order", f.OrderCreate)
	e.GET("/library/order/:o_id", f.FindOrderById)
	e.DELETE("/library/order/:o_id", f.DelOrderById)

	e.Static("/", "C:/Users/박수빈/Desktop/Library-API/front/monitorapp-library-react-app/build/index.html")

	//start server
	e.Start(":5000")

}
