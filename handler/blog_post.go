package handler

import (
	"fmt"
	"html"
	"os"

	"github.com/Michelsen93/my-blog/view/blog_post"
	"github.com/labstack/echo/v4"
	"github.com/russross/blackfriday"
)

type BlogPostHandler struct {
}

const postsPath = "posts/";

func (h BlogPostHandler) HandleBlogPostShow(c echo.Context) error {
    return render(c, blog_post.Show(fileToHtml()))
}

func fileToHtml() string{
    b, err := os.ReadFile(fmt.Sprintf("%vhello_world.md", postsPath))
    if err != nil {
        fmt.Print(err)
    }
    r := blackfriday.MarkdownCommon(b)
    return html.UnescapeString(string(r))
}
