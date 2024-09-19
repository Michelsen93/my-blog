package main

    import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Michelsen93/my-blog/handler"
	"github.com/labstack/echo/v4"
)
func main() {
    app := echo.New()
    app.Static("/", "assets")
    indexHandler := handler.IndexHandler{}
    blogpostHandler := handler.BlogPostHandler{}
    blogListHanler := handler.BlogListHandler{}
    app.GET("/", indexHandler.HandleIndexShow)
    app.GET("/blog/:path", blogpostHandler.HandleBlogPostShow)
    app.GET("/blog_list", blogListHanler.HandleBlogListShow)
    ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
    defer stop()
    port := os.Getenv("HTTP_PLATFORM_PORT");
    if port == "" {
        port = "8080"
    }
    go func() {
		if err := app.Start(port); err != nil && err != http.ErrServerClosed {
			app.Logger.Fatal("shutting down the server")
		}
        fmt.Println("Server running")
	}()
    <-ctx.Done()
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
	if err := app.Shutdown(ctx); err != nil {
		app.Logger.Fatal(err)
	}
}
