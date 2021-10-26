package main

import (
	"Library_api/crawler"
	"Library_api/f"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORS())
	//writing image sha256:9e385597603dd5a4838da36c2618a5cfb7b19d799ff486fa59980f9ed46d7828
	// Routes
	e.POST("/", f.Signup)
	e.POST("/login", f.Login)
	e.GET("/logout", f.Logout)

	e.POST("/book", f.BookCreate)
	e.GET("/book/code", f.FindBookByStatus)
	e.GET("/book", f.FindBookByKeyword)
	e.PUT("/book/:keyword", f.BookUpdate)
	e.DELETE("/book/:code4:", f.DelBookByStatusCode)
	//del keyword => id 변경

	e.POST("/user", f.UserCreate)
	e.GET("/user/:keyword", f.FindUserByKeyword)
	e.GET("/user/code/:keyword", f.FindUserByStatusCode)
	e.PUT("/user/:keyword", f.UserUpdate)
	e.DELETE("/user/key", f.DelUserByKeyword)
	e.DELETE("/user/code3", f.DelUserByStatusCode)

	e.POST("/library/rent", f.Rent)
	e.POST("/library/return", f.Return)

	//# open API
	e.GET("/other/yes24", crawler.Crawling)
	e.GET("/other/kyobo", crawler.Crawling)
	//# Crawler
	e.GET("/other/interpark", crawler.Interpark)
	e.GET("/other/kakao", crawler.Kakao)

	//start server
	e.Start(":4000")

}

//우분투 20.04 -> 도커 파일 -> 도커 라이즈
