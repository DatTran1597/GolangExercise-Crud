package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func createUser(c echo.Context) error {
	return c.JSON(http.StatusCreated, u)
}

func main() {
	e := echo.New()
	p := e.Group("/users")
	p.POST("/", createUser)
	e.Start(":1323")
}
