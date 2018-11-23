package main

import "github.com/labstack/echo"

func main() {
	p:= e.Group("/users")
	p.POST("")
	e := echo.New()
	e.Start(":1323")
}
