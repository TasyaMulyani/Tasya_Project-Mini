package controller

import (
	"net/http"

	"Tasya_Project-Mini/model"
	"Tasya_Project-Mini/usecase"

	"github.com/labstack/echo/v4"
)

func LoginUserController(c echo.Context) error {
	user := model.User{}
	c.Bind(&user)

	err := usecase.LoginUser(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success login",
		"user":   user,
	})
}
