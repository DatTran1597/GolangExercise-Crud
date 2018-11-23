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

func show(h echo.Context) error {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	coll := session.DB("mydb").C("demo")

	// Find the number of games won by Dave
	player := "Dave"
	gamesWon, err := coll.Find(bson.M{"winner": player}).Count()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s has won %d games.\n", player, gamesWon)
	var result Game
	err = coll.Find(bson.M{"winner": player, "location": "Austin"}).Select(bson.M{"official_game": 1}).One(&result)
	if err != nil {
		panic(err)
	}
	fmt.Println("Is game in Austin Official?", result.OfficialGame)
	return h.String(http.StatusOK, player+" has won "+strconv.Itoa(gamesWon)+" games\n")
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	p := e.Group("/users")
	p.POST("/", createUser)
	p.GET("/:id", getUser)
	e.Start(":1323")
}
