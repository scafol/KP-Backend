package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/scafol/KP-Backend/entity"
	"github.com/scafol/KP-Backend/model"
	"github.com/scafol/KP-Backend/utils"
)

type CategoryController struct {
	model model.CategoryModel
}

// InsertCategory - create new category
// POST /Category
// Create a new Category
func (CategoryController CategoryController) InsertCategory(c echo.Context) error {
	var input entity.Category

	if err := c.Bind(&input); err != nil {
		res := utils.MakeResponse("Failed", err.Error(), nil)
		return c.JSON(http.StatusBadRequest, res)
	}

	category, insertError := CategoryController.model.InsertCategory(input)
	if insertError != nil {
		res := utils.MakeResponse("Failed", insertError.Error(), nil)
		return c.JSON(http.StatusBadRequest, res)
	}

	res := utils.MakeResponse("Success", "New category successfully inserted", &category)
	return c.JSON(http.StatusBadRequest, res)
}

// FindCategories - retrieve all category
// GET /categories
// Find all users
func (CategoryController CategoryController) FindCategories(c echo.Context) error {
	categories := CategoryController.model.GetCategories()
	res := utils.MakeResponse("Success", "Categories successfully loaded", categories)
	return c.JSON(http.StatusOK, res)
}

// FindCategory - retrevie discussion by id
// GET /category/{id}
// Find specific category base on id
func (CategoryController CategoryController) FindCategory(c echo.Context) error {
	id := c.Param("id")
	category := CategoryController.model.GetCategory(id)
	res := utils.MakeResponse("Success", "Category Successfully loaded", category)
	return c.JSON(http.StatusOK, res)
}

// UpdateCategory - Update Category
// PUT /category/{id}
// Update Category base on id
func (CategoryController CategoryController) UpdateCategory(c echo.Context) error {
	var input entity.Category

	id := c.Param("id")

	if err := c.Bind(&input); err != nil {
		res := utils.MakeResponse("Failed", err.Error(), nil)
		return c.JSON(http.StatusBadRequest, res)
	}

	category, err := CategoryController.model.UpdateCategory(input, id)
	if err != nil {
		res := utils.MakeResponse("Failed", err.Error(), nil)
		return c.JSON(http.StatusBadRequest, res)
	}

	res := utils.MakeResponse("Success", "Category successfully updated", category)
	return c.JSON(http.StatusCreated, res)
}

// DeleteCategory - Delete a Category
// DELETE /category/{id}
// Delete a Category base on id
func (CategoryController CategoryController) DeleteCategory(c echo.Context) error {
	id := c.Param("id")

	err := CategoryController.model.DeleteCategory(id)
	if err != nil {
		res := utils.MakeResponse("Failed", err.Error(), nil)
		return c.JSON(http.StatusBadGateway, res)
	}

	res := utils.MakeResponse("Success", "Category successfully deleted", nil)
	return c.JSON(http.StatusOK, res)
}
