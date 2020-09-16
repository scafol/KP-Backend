package routes

import (
	"Golang-Echo-MVC-Pattern/controller"
	"github.com/labstack/echo"
)

type Routing struct {
	example controller.ExampleController
	user controller.UserController
	catagory controller.CatagoryController
	discussion controller.DiscussionController
}

func (Routing Routing) GetRoutes() *echo.Echo {
	e := echo.New()
	e.GET("/posts/", Routing.example.GetPostsController)

	// users
	e.POST("/user/", Routing.user.AddUserController)
	e.GET("/user/", Routing.user.ViewUsersController)
	e.GET("/user/:idUser", Routing.user.ViewUserIdController)

	// catagory
	e.POST("/catagory/", Routing.catagory.AddCatagory)
	e.GET("/catagory/", Routing.catagory.ViewAllCatagory)

	// Discussion
	e.POST("/discussion/", Routing.discussion.AddDiscussion)

	return e
}
