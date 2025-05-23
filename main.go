package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/smnschmnck/roundest-go-htmx/db"
	"github.com/smnschmnck/roundest-go-htmx/pages"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	err = db.InitDb()
	if err != nil {
		panic(err.Error())
	}

	e := echo.New()
	e.Use(middleware.Gzip())

	e.Static("/static", "./.build/static")

	pages.RegisterPages(e)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
