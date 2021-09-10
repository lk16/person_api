package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func helloHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, Rigon!")
}

func main() {
	e := echo.New()

	e.GET("/", helloHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
