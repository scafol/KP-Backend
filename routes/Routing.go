package routes

import (
	"github.com/labstack/echo"
	"github.com/scafol/KP-Backend/controller"
)

type Routing struct {
	example    controller.ExampleController
	user       controller.UserController
	category   controller.CategoryController
	discussion controller.DiscussionController
	comment    controller.CommentController
	upload     controller.UploaderController
}

func (Routing Routing) GetRoutes() *echo.Echo {
	e := echo.New()

	e.GET("/posts/", Routing.example.GetPostsController)

	// user context
	e.POST("/user", Routing.user.CreateUser)
	e.GET("/users", Routing.user.GetUsers)
	e.GET("/user/:id", Routing.user.GetUser)

	// category context
	e.GET("/categories", Routing.category.FindCategories)
	e.GET("/category/:id", Routing.category.FindCategory)
	e.POST("/category", Routing.category.InsertCategory)
	e.PUT("/category/:id", Routing.category.EditCategory)
	e.DELETE("/category/:id", Routing.category.DeleteCategory)

	// discussion context
	e.GET("/discussions", Routing.discussion.FindDiscussions)
	e.GET("/discussion/:id", Routing.discussion.FindDiscussion)
	e.POST("/discussion", Routing.discussion.InsertDiscussion)
	e.PUT("/discussion/:id", Routing.discussion.EditDiscussion)
	e.DELETE("/discussion/:id", Routing.discussion.DeleteDiscussion)

	// comment context
	e.GET("/comments", Routing.comment.FindComments)
	e.GET("/comment/:id", Routing.comment.FindComment)
	e.POST("/comment", Routing.comment.InsertComment)
	e.PUT("/comment/:id", Routing.comment.EditComment)
	e.DELETE("/comment/:id", Routing.comment.DeleteComment)

	// uploader service
	e.POST("/upload", Routing.upload.Uploader)
	e.POST("/uploads", Routing.upload.Uploaders)
	return e
}
