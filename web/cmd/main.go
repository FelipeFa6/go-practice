package main

import (
	"github.com/felipefa6/todo-go/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	th := handler.TaskHandler{}
	app.GET("/", th.HandleTaskShow)

	app.Logger.Fatal(app.Start(":3000"))
}
