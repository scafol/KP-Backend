package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/scafol/KP-Backend/model"
	"github.com/scafol/KP-Backend/responsegraph"
)

type ExampleController struct {
	model model.ExampleModel
}

// Get Example Controller
func (ExampleController ExampleController) GetPostsController(c echo.Context) error {
	posts := ExampleController.model.GetPosts()
	res := responsegraph.ResponseGeneric{
		Status:  "Success",
		Message: "Posts Loaded",
		Data:    posts,
	}
	return c.JSON(http.StatusOK, res)
}
