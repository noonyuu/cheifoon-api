package handler

import (
	"myapp/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SetMachinePOST(c echo.Context) error {

	machine := new(model.Machine)
	if err := c.Bind(machine); err != nil {
		c.Logger().Error(err)
		return err
	}

	err := model.CreateMachine(machine)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "machine creation failed",
		}
	}

	return c.JSON(http.StatusCreated, machine)
}
