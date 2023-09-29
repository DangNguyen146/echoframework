package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func HandleSignIn(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"user":  "1",
		"email": "1@gmail.com",
	})
}

func HandleSignUp(c echo.Context) error {
	type User struct {
		Email    string
		FullName string
		Age      int
	}
	user := User{
		Email:    "1@gmail.com",
		FullName: "2",
		Age:      1,
	}
	return c.JSON(http.StatusOK, user)
}
