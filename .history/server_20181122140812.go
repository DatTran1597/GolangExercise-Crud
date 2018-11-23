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
var seq=1

func createUser(c echo.Context) error {
	u:=$user{
		ID:seq
	}
	if err:=c.Bind(u); err!=nil{
		panic(err)
	}
	users[u.ID]:=u
	seq++
	return c.JSON(http.StatusCreated, u)
}

func main() {
	e := echo.New()
	p := e.Group("/users")
	p.POST("/", createUser)
	e.Start(":1323")
}
