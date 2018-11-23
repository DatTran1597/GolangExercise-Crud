package main

import (
	"fmt"
	"net/http"
	"time"

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

type Player struct {
	Name   string    `bson:"name"`
	Decks  [2]string `bson:"decks"`
	Points uint8     `bson:"points"`
	Place  uint8     `bson:"place"`
}
type Game struct {
	Winner       string    `bson:"winner"`
	OfficialGame bool      `bson:"official_game"`
	Location     string    `bson:"location"`
	StartTime    time.Time `bson:"start"`
	EndTime      time.Time `bson:"end"`
	Players      []Player  `bson:"players"`
}

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
	//var result Game
	err := coll.Find(bson.M{"_id": "1"}).Select(bson.M{"name": "11"}).One(&result)

	//err = coll.Find(bson.M{"winner": player, "location": "Austin"}).Select(bson.M{"official_game": 1}).One(&result)

	//var result []struct{ Value int }
	if abc != nil {
		panic(err)
	}

	fmt.Println(result)
	return c.JSON(http.StatusOK, result)
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	p := e.Group("/users")
	p.POST("/", createUser)
	p.GET("/:id", getUser)
	e.Start(":1323")
}
