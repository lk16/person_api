package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"

	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	ID   string `json:"id"`
}

// var personDB map[int]Person

var dbClient *mongo.Client

/* func listPersons(c echo.Context) error {
	personsList := []Person{}

	for _, person := range personDB {
		personsList = append(personsList, person)
	}

	return c.JSON(http.StatusOK, personsList)
} */

func createPerson(c echo.Context) error {
	var person Person

	if err := c.Bind(&person); err != nil {
		return err
	}

	// personDB[id] = person

	collection := dbClient.Database("demo").Collection("person")

	result, err := collection.InsertOne(context.Background(), person)

	if err != nil {
		return c.String(http.StatusInternalServerError, "Database error")
	}

	person.ID = result.InsertedID.(primitive.ObjectID).Hex()

	return c.JSON(http.StatusCreated, person)
}

/* func getPerson(c echo.Context) error {
	idString := c.Param("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "id parameter is not an integer")
	}

	person, ok := personDB[id]
	if !ok {
		return c.JSON(http.StatusNotFound, "Could not find person")
	}

	return c.JSON(http.StatusOK, person)
}

func deletePerson(c echo.Context) error {
	idString := c.Param("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "id parameter is not an integer")
	}

	_, ok := personDB[id]
	if !ok {
		return c.JSON(http.StatusNotFound, "Could not find person")
	}

	delete(personDB, id)
	return c.NoContent(http.StatusOK)
} */

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Printf("Failed to connect to DB.")
		os.Exit(1)
	}

	dbClient = client

	//personDB = make(map[int]Person)
	e := echo.New()
	// e.GET("/greet/:user", greetingHandler)
	// e.GET("/persons", listPersons)
	e.POST("/persons", createPerson)
	// e.GET("/persons/:id", getPerson)
	// e.DELETE("/persons/:id", deletePerson)

	e.Logger.Fatal(e.Start(":1323"))
}
