package main

import (
	"jar-project/database"
	"jar-project/routes"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database.Database()
	e := routes.Routes()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Logger.Fatal(e.Start(":1324"))
}
