package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	ID   int    `json:"id"`
}

var personDB map[int]Person

func listPersons(c echo.Context) error {
	personsList := []Person{}

	for _, person := range personDB {
		personsList = append(personsList, person)
	}

	return c.JSON(http.StatusOK, personsList)
}

func createPerson(c echo.Context) error {
	var person Person

	if err := c.Bind(&person); err != nil {
		return err
	}

	id := rand.Intn(100)
	person.ID = id

	personDB[id] = person
	return c.JSON(http.StatusOK, person)
}

func getPerson(c echo.Context) error {
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
}

func greetingHandler(c echo.Context) error {
	user := c.Param("user")
	response := fmt.Sprintf("Hello, %s\n", user)
	return c.String(http.StatusOK, response)
}

func main() {
	personDB = make(map[int]Person)
	e := echo.New()
	e.GET("/greet/:user", greetingHandler)
	e.GET("/persons", listPersons)
	e.POST("/persons", createPerson)
	e.GET("/persons/:id", getPerson)
	e.DELETE("/persons/:id", deletePerson)

	e.Logger.Fatal(e.Start(":1323"))
}
