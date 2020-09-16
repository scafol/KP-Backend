package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/scafol/KP-Backend/entity"
	"github.com/scafol/KP-Backend/model"
	"github.com/scafol/KP-Backend/utils"
)

type UserController struct {
	model model.UserModel
}

func (UserController UserController) CreateUser(c echo.Context) error {
	var input entity.User

	if err := c.Bind(&input); err != nil {
		res := utils.MakeResponse("Failed", err.Error(), nil)
		return c.JSON(http.StatusBadRequest, res)
	}

	user, insertError := UserController.model.InsertUser(&input)
	if insertError != nil {
		res := utils.MakeResponse("Failed", insertError.Error(), nil)
		return c.JSON(http.StatusBadRequest, res)
	}

	res := utils.MakeResponse("Success", "New user successfully inserted", user)
	return c.JSON(http.StatusBadRequest, res)
}

func (UserController UserController) GetUsers(c echo.Context) error {
	users := UserController.model.GetUsers()
	res := utils.MakeResponse("Successfully", "Data successfully loaded", users)
	return c.JSON(http.StatusOK, res)
}

func (UserController UserController) GetUser(c echo.Context) error {
	id := c.Param("id")
	user := UserController.model.GetUser(id)
	res := utils.MakeResponse("Successfully", "Data successfully loaded", user)
	return c.JSON(http.StatusOK, res)
}
