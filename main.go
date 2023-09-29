package main

import (
	"echoframework/db"
	"echoframework/handler"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"log"
	"os"
)

func main() {
	env := godotenv.Load()
	if env != nil {
		log.Fatalf("Some error occured. Err: %s", env)
	}

	sql := &db.Sql{
		Host:     os.Getenv("HOST_DB"),
		Port:     os.Getenv("PORT_DB"),
		UserName: os.Getenv("USERNAME_DB"),
		Password: os.Getenv("PASSWORD_DB"),
		DbName:   os.Getenv("DBNAME_DB"),
	}
	sql.Connect()
	defer sql.Close()

	e := echo.New()
	e.GET("/", handler.Welcome)
	e.Logger.Fatal(e.Start(":3000"))
}
