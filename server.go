package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(context echo.Context) error {
		return context.JSON(http.StatusOK,"Hello World")
	})
	e.POST("/user", CreateUser)

	e.Start(":1323")
}

type User struct {
	Name string `json:"name" query:"name"`
	Age int8 `json:"age" query:"age"`
}

func CreateUser(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusOK,u.Name)
}
