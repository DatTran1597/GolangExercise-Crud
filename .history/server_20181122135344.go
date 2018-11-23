package main

import "github.com/labstack/echo"

func main() {
	e := echo.New()
	p := e.Group("/users")
	p.POST("/")
	e.Start(":1323")
}
