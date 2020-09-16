package controller

import (
	"Golang-Echo-MVC-Pattern/entity"
	"Golang-Echo-MVC-Pattern/model"
	"Golang-Echo-MVC-Pattern/utils"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type DiscussionController struct {
	discussionModel model.DiscussionModel
	userModel       model.UserModel
	catagoryModel   model.CatagoryModel
}

func (e *DiscussionController) AddDiscussion(c echo.Context) error {
	var discussion = entity.Discussion{}
	err := c.Bind(&discussion)
	if err != nil {
		fmt.Printf("[DiscussionController.AddDiscussion] error bind data %v \n", err)
		return utils.HandleError(c, http.StatusInternalServerError, err.Error())
	}
	_, err = e.catagoryModel.GetCatagoryId(int(discussion.CatagoryID))
	if err != nil {
		return utils.HandleError(c, http.StatusNotFound, err.Error())
	}
	_, err = e.userModel.GetUserById(int(discussion.UserID))
	if err != nil {
		return utils.HandleError(c, http.StatusNotFound, err.Error())
	}
	if discussion.Message == "" || discussion.Title == "" {
		return utils.HandleError(c, http.StatusBadRequest, "field are required")
	}
	mDiscussion, err := e.discussionModel.AddDiscussion(&discussion)
	if err != nil {
		return utils.HandleError(c, http.StatusInternalServerError, err.Error())
	}
	return utils.HandleSuccess(c, mDiscussion)
}
