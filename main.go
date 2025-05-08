package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/smnschmnck/roundest-go-htmx/db"
	"github.com/smnschmnck/roundest-go-htmx/pages"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = db.InitDb()
	if err != nil {
		panic(err.Error())
	}

	e := echo.New()
	e.Static("/static", "./.build/static")

	pages.RegisterPages(e)

	e.Logger.Fatal(e.Start(":1323"))
}
