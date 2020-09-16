package routes

import (
	"github.com/labstack/echo"
	"github.com/scafol/KP-Backend/controller"
)

type Routing struct {
	example  controller.ExampleController
	user     controller.UserController
	category controller.CategoryController
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
	e.PUT("/category/:id", Routing.category.UpdateCategory)
	e.DELETE("/category/:id", Routing.category.DeleteCategory)
	return e
}
