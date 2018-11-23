package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type user struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = map[int]*user{}

func createUser(c echo.Context) error {

	return c.JSON(http.StatusCreated, u)
}

func main() {
	e := echo.New()
	p := e.Group("/users")
	p.POST("/", createUser)
	e.Start(":1323")
}
