package handler

import (
	"github.com/Michelsen93/my-blog/view/blog_post"
	"github.com/labstack/echo/v4"
)

type BlogPostHandler struct {

}

func (h BlogPostHandler) HandleBlogPostShow(c echo.Context) error {
    return render(c, blog_post.Show("<p>hello<p>"))
}
