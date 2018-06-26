package main

import (
	"flag"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//go:generate go-assets-builder -s="/static" -o assets.go static/

var (
	isProd bool
	db     *gorm.DB
)

func init() {
	flag.BoolVar(&isProd, "prod", false, "print version")
}

func main() {
	flag.Parse()

	e := echo.New()

	if isProd {
		fmt.Println("serve from Assets")
		e.GET("/*", echo.WrapHandler(http.FileServer(Assets)))
	} else {
		fmt.Println("serve from static dir")
		e.Static("/", "static")
	}

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/api/reservations", get)

	e.Logger.Fatal(e.Start(":8080"))
}

func debug() []Reservation {
	reservations := []Reservation{
		Reservation{Date: "06/26/2018", Name: "himetani"},
		Reservation{Date: "06/27/2018", Name: "taro"},
	}

	return reservations

}

type Reservation struct {
	Date string `json:"date"`
	Name string `json:"name"`
}

func get(c echo.Context) error {
	return c.JSON(http.StatusOK, debug())
}

func create(c echo.Context) error {

	return nil
}
