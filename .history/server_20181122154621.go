package main

import (
	"net/http"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2"
)

type user struct {
	ID   int    `bson:"_id"`
	Name string `bson:"name"`
}

var users = map[int]*user{}
var seq = 1

func createUser(c echo.Context) error {
	u := &user{
		ID: seq,
	}
	if err := c.Bind(u); err != nil {
		panic(err)
	}
	users[u.ID] = u
	seq++
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	coll := session.DB("mydb").C("demo")
	if err := coll.Insert(u); err != nil {
		panic(err)
	}
	return c.JSON(http.StatusCreated, u)
}

func getUser(c echo.Context) error {

}

func main() {
	e := echo.New()
	p := e.Group("/users")
	p.POST("/", createUser)
	p.GET("/:id", getUser)
	e.Start(":1323")
}
