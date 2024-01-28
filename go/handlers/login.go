package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/muling3/redis/go/views"
)

func Login(c echo.Context) error {
	return views.Login().Render(c.Request().Context(), c.Response())
}
