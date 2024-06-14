package controller

import (
	"github.com/labstack/echo/v4"
)

type controller struct {
}

func NewControler() controller {
	return controller{}
}

func (c *controller) Home(e echo.Context) error {

	return e.JSON(202, "Hello welcome, your access is free")
}
