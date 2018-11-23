package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID   int    `bson:"_id"`
	Name string `bson:"name"`
}

var users = map[int]*User{}
var seq = 1

func createUser(c echo.Context) error {
	u := &User{
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
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	coll := session.DB("mydb").C("demo")

	var result User
	err = coll.Find(nil).Select(bson.M{"name": "vegeta"}).One(&result)

	//err = coll.Find(bson.M{"winner": player, "location": "Austin"}).Select(bson.M{"official_game": 1}).One(&result)

	//var result []struct{ Value int }
	if err != nil {
		panic(err)
	}

	fmt.Println("result")
	return c.JSON(http.StatusOK, result)
}

func updateUser(c echo.Context) error {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	coll := session.DB("mydb").C("demo")
	selector := bson.M{"name": "goku god"}
	updator := bson.M{"$set": bson.M{"name": "goku"}}
	info, err := coll.UpdateAll(selector, updator)
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK, info.Updated)
}

func deleteUser(c echo.Context) error {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	coll := session.DB("mydb").C("demo")
	info, err := coll.RemoveAll(bson.M{"name": c.Param("id")})
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK, info.Removed)
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	p := e.Group("/users")
	p.POST("/", createUser)
	p.GET("/:id", getUser)
	p.PUT("/:id", updateUser)
	p.DELETE("/:id", deleteUser)
	e.Start(":1323")
}
