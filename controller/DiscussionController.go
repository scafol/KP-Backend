package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/scafol/KP-Backend/entity"
	"github.com/scafol/KP-Backend/model"
	"github.com/scafol/KP-Backend/utils"
)

type DiscussionController struct {
	model model.DiscussionModel
}

// InsertDiscussion - create new discussion
// POST /discussion
// Create a new discussion
func (discussionController DiscussionController) InsertDiscussion(c echo.Context) error {
	var input entity.Discussion

	if err := c.Bind(&input); err != nil {
		res := utils.MakeResponse("Failed", err.Error(), nil)
		return c.JSON(http.StatusBadRequest, res)
	}

	discussion, insertError := discussionController.model.CreateDiscussion(&input)
	if insertError != nil {
		res := utils.MakeResponse("Failed", insertError.Error(), nil)
		return c.JSON(http.StatusBadRequest, res)
	}

	res := utils.MakeResponse("Success", "New discussion successfully inserted", &discussion)
	return c.JSON(http.StatusBadRequest, res)
}

// FindDiscussions - retrieve all category
// GET /Discussions
// Find all discussions
func (discussionController DiscussionController) FindDiscussions(c echo.Context) error {
	discussions, err := discussionController.model.GetDiscussions()
	if err != nil {
		res := utils.MakeResponse("Failed", err.Error(), nil)
		return c.JSON(http.StatusOK, res)
	}
	// final discussions data
	discussionsData := []map[string]interface{}{}

	for _, discussion := range discussions {
		discussionsData = append(discussionsData, map[string]interface{}{
			"Data":         discussion,
			"CommentCount": len(discussion.Comments),
		})
	}
	res := utils.MakeResponse("Success", "Discussions successfully loaded", discussionsData)
	return c.JSON(http.StatusOK, res)
}

// FindDiscussion - retrevie discussion by id
// GET /discussion/{id}
// Find specific discussion base on id
func (discussionController DiscussionController) FindDiscussion(c echo.Context) error {
	id := c.Param("id")
	discussion, err := discussionController.model.GetDiscussion(id)
	if err != nil {
		res := utils.MakeResponse("Failed", err.Error(), nil)
		return c.JSON(http.StatusOK, res)
	}
	res := utils.MakeResponse("Success", "Discussion Successfully loaded", &discussion)
	return c.JSON(http.StatusOK, res)
}

// EditDiscussion - Edit Discussion
// PUT /discussion/{id}
// Update discussion base on id
func (discussionController DiscussionController) EditDiscussion(c echo.Context) error {
	var input entity.Discussion

	id := c.Param("id")

	if err := c.Bind(&input); err != nil {
		res := utils.MakeResponse("Failed", err.Error(), nil)
		return c.JSON(http.StatusBadRequest, res)
	}

	discussion, err := discussionController.model.UpdateDiscussion(&input, id)
	if err != nil {
		res := utils.MakeResponse("Failed", err.Error(), nil)
		return c.JSON(http.StatusBadRequest, res)
	}

	res := utils.MakeResponse("Success", "Discussion successfully updated", &discussion)
	return c.JSON(http.StatusCreated, res)
}

// DeleteDiscussion - Delete a iscussion
// DELETE /discussion/{id}
// Delete a discussion base on id
func (discussionController DiscussionController) DeleteDiscussion(c echo.Context) error {
	id := c.Param("id")

	err := discussionController.model.DeleteDiscussion(id)
	if err != nil {
		res := utils.MakeResponse("Failed", err.Error(), nil)
		return c.JSON(http.StatusBadGateway, res)
	}

	res := utils.MakeResponse("Success", "Discussion successfully deleted", nil)
	return c.JSON(http.StatusOK, res)
}
