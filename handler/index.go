package handler

import (
    "github.com/labstack/echo/v4"
    "github.com/Michelsen93/my-blog/view/index"
)

type IndexHandler struct {

}

func (h IndexHandler) HandleIndexShow(c echo.Context) error {
    return render(c, index.Show())
}
