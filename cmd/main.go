package main

import (
	"fmt"
	"github.com/Michelsen93/my-blog/handler"
	"github.com/labstack/echo/v4"
)
func main() {
    app := echo.New()
    app.Static("/", "assets")
    indexHandler := handler.IndexHandler{}
    blogpostHandler := handler.BlogPostHandler{}
    app.GET("/", indexHandler.HandleIndexShow)
    app.GET("/blog", blogpostHandler.HandleBlogPostShow)

    app.Start(":6969")
    fmt.Println("Server running")
}

