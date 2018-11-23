package main

import (
	"fmt"
	"net/http"
	"strconv"

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
	idget, err := strconv.Atoi(c.Param("id"))
	err = coll.Find(bson.M{"_id": idget}).Select(bson.M{"name": idget}).One(&result)

	//err = coll.Find(bson.M{"winner": player, "location": "Austin"}).Select(bson.M{"official_game": 1}).One(&result)

	//var result []struct{ Value int }
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
	return c.JSON(http.StatusOK, result)
}

func updateUser(c echo.Context) error {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	coll := session.DB("mydb").C("demo")
	idup, err := strconv.Atoi(c.Param("id"))
	selector := bson.M{"_id": idup}
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
	iddel, err := strconv.Atoi(c.Param("id"))
	err = coll.Remove(bson.M{"_id": iddel})
	if err != nil {
		panic(err)
	}
	//fmt.Println(info.Removed)
	return c.JSON(http.StatusOK, err)
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
