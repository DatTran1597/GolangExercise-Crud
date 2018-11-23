package main

import "github.com/labstack/echo"

func createUser(c echo.Context) error {

}

func main() {
	e := echo.New()
	p := e.Group("/users")
	p.POST("/", createUser)
	e.Start(":1323")
}
