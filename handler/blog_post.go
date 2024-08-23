package handler

import (
	"fmt"
	"html"
	"os"

	"github.com/Michelsen93/my-blog/view/blog_post"
	"github.com/labstack/echo/v4"
	"github.com/russross/blackfriday"
)

type BlogPostHandler struct {}
type BlogPostParams struct {
    Path string `param:"path"`
}


const postsPath = "posts/";

func (h BlogPostHandler) HandleBlogPostShow(c echo.Context) error {
    params := new(BlogPostParams)
    if err := c.Bind(params); err != nil {
        //Could not find blog post
        return err
    }
    return render(c, blog_post.Show(fileToHtml(params.Path)))
}

func fileToHtml(path string) string{
    b, err := os.ReadFile(fmt.Sprintf("%v%v", postsPath, path))
    if err != nil {
        // TODO - add errorpage
        return "Could not find this blog post"
    }
    r := blackfriday.MarkdownCommon(b)
    return html.UnescapeString(string(r))
}
