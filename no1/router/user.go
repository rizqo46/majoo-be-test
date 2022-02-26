package router

import (
	"majoo-be-test/controller"

	"github.com/labstack/echo/v4"
)

type UserRouter struct {
	userController controller.UserController
}

func NewUserRouter(userController controller.UserController) UserRouter {
	return UserRouter{userController}
}

func (r *UserRouter) Mount(group *echo.Group) {
	group.POST("/login", r.userController.Login)
	group.POST("/logout", r.userController.Logout)
}
