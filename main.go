package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

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

	var err error
	db, err = gorm.Open("mysql", "coffee:aeropress@tcp(localhost:3306)/coffeetime?parseTime=true")
	if err != nil {
		panic(err)
	}

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
	e.GET("/api/debug", create)
	e.PUT("/api/reservations", create)

	e.Logger.Fatal(e.Start(":8080"))
}

func debug() []Reservation {
	t0, _ := time.Parse("2006-01-02", "2018-06-26")
	t1, _ := time.Parse("2006-01-02", "2018-06-27")
	reservations := []Reservation{
		Reservation{Date: t0, Name: "himetani"},
		Reservation{Date: t1, Name: "taro"},
	}
	return reservations
}

type Reservation struct {
	gorm.Model
	Date time.Time `json:"date"`
	Name string    `json:"name"`
}

type response struct {
	ExpiredReservations []Reservation `json:"expired_reservations"`
	ValidReservations   []Reservation `json:"valid_reservations"`
}

func get(c echo.Context) error {
	reservations := []Reservation{}
	db.Find(&reservations)

	yesterday := time.Now().Add(24 * -time.Hour)

	var idx int
	for i, r := range reservations {
		if r.Date.After(yesterday) {
			idx = i
			break
		}
	}

	res := &response{
		ExpiredReservations: reservations[0:idx],
		ValidReservations:   reservations[idx:],
	}
	return c.JSON(http.StatusOK, res)
}

func create(c echo.Context) error {
	name := c.FormValue("name")
	date := c.FormValue("date")

	t, err := time.Parse(time.RFC3339, date)
	if err != nil {
		return err
	}

	db.Create(&Reservation{Date: t, Name: name})
	return nil
}
