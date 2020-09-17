package controller

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/scafol/KP-Backend/entity"
	"github.com/scafol/KP-Backend/model"
	"github.com/scafol/KP-Backend/utils"
)

type CommentController struct {
	db model.CommentModel
}

// InsertComment - insert new comment
// POST /comment
// Insert new comment related to a Comment
func (commentController CommentController) InsertComment(c echo.Context) error {
	var input entity.Comment

	if err := c.Bind(&input); err != nil {
		res := utils.MakeResponse("Failed", err.Error(), nil)
		return c.JSON(http.StatusBadRequest, res)
	}

	// insert into database
	comment, err := commentController.db.CreateComment(&input)
	if err != nil {
		res := utils.MakeResponse("Failed", err.Error(), nil)
		return c.JSON(http.StatusBadGateway, res)
	}

	res := utils.MakeResponse("Success", "Comment successfully created", &comment)
	return c.JSON(http.StatusCreated, res)
}

// FindComments - retrieve all comments
// GET /comments
// Find all comments
func (commentController CommentController) FindComments(c echo.Context) error {
	comments, err := commentController.db.FindComments()

	if err != nil {
		res := utils.MakeResponse("Failed", err.Error(), nil)
		return c.JSON(http.StatusInternalServerError, res)
	}
	res := utils.MakeResponse("Success", "Comments successfully loaded", comments)
	return c.JSON(http.StatusOK, res)
}

// FindComment - retrieve all comments
// GET /comment
// Find all comment
func (commentController CommentController) FindComment(c echo.Context) error {
	id := c.Param("id")

	if id == "" {
		res := utils.MakeResponse("Failed", "id must not be empty", nil)
		return c.JSON(http.StatusBadGateway, res)
	}

	comment, err := commentController.db.FindComment(id)
	if err != nil {
		res := utils.MakeResponse("Failed", err.Error(), nil)
		return c.JSON(http.StatusBadGateway, res)
	}

	if comment.ID == 0 {
		errorMessage := fmt.Sprintf("Comment with %d id not found", comment.ID)
		res := utils.MakeResponse("Failed", errorMessage, nil)
		return c.JSON(http.StatusNotFound, res)
	}

	res := utils.MakeResponse("Success", "Comment successfully loaded", &comment)
	return c.JSON(http.StatusOK, res)
}

// EditComment - Update comment
// PUT /comment/{id}
// Update comment base on id
func (commentController CommentController) EditComment(c echo.Context) error {
	var input entity.Comment

	id := c.Param("id")

	if err := c.Bind(&input); err != nil {
		res := utils.MakeResponse("Failed", err.Error(), nil)
		return c.JSON(http.StatusBadRequest, res)
	}

	comment, err := commentController.db.UpdateComment(&input, id)
	if err != nil {
		res := utils.MakeResponse("Failed", err.Error(), nil)
		return c.JSON(http.StatusBadRequest, res)
	}
	res := utils.MakeResponse("Success", "Comment successfully updated", &comment)
	return c.JSON(http.StatusOK, res)
}

// DeleteComment - Delete a Comment
// DELETE /Comment/{id}
// Delete a Comment base on id
func (commentController CommentController) DeleteComment(c echo.Context) error {
	id := c.Param("id")

	err := commentController.db.DeleteComment(id)
	if err != nil {
		res := utils.MakeResponse("Failed", err.Error(), nil)
		return c.JSON(http.StatusBadRequest, res)
	}
	res := utils.MakeResponse("Success", "Comment successfully deleted", nil)
	return c.JSON(http.StatusOK, res)
}
