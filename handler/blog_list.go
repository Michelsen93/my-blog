package handler

import (
	"github.com/Michelsen93/my-blog/view/blog_list"
	"github.com/labstack/echo/v4"
)



type BlogListHandler struct {

}

func (h BlogListHandler) HandleBlogListShow(c echo.Context) error {
    return render(c, blog_list.Show())
}
