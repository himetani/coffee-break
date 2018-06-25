package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//go:generate go-assets-builder -s="/static" -o assets.go static/

var isProd bool

func init() {
	flag.BoolVar(&isProd, "prod", false, "print version")
}

func main() {
	flag.Parse()

	e := echo.New()

	if isProd {
		fmt.Println("serve from Assets")
		e.GET("/", echo.WrapHandler(http.FileServer(Assets)))
	} else {
		fmt.Println("serve from static dir")
		e.GET("/", echo.WrapHandler(http.FileServer(http.Dir("./static"))))
	}

	e.Use(middleware.CORS())
	e.GET("/api/reservations", get)

	e.Logger.Fatal(e.Start(":8080"))
}

func debug() []reservation {
	reservations := []reservation{
		reservation{"06/26/2018", "himetani"},
		reservation{"06/27/2018", "taro"},
	}

	return reservations

}

type reservation struct {
	Date string `json:"date"`
	Name string `json:"name"`
}

func get(c echo.Context) error {
	return c.JSON(http.StatusOK, debug())
}

func create(c echo.Context) error {
	return nil
}
