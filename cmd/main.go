package main

import (
	"fmt"

	"github.com/Michelsen93/my-blog/handler"
	"github.com/labstack/echo/v4"
)
func main() {
    app := echo.New()
    indexHandler := handler.IndexHandler{}
    app.GET("/", indexHandler.HandleIndexShow)

    app.Start(":6969")
    fmt.Println("Server running")
}

