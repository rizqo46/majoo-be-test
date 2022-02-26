package controller

import (
	"majoo-be-test/config"
	"majoo-be-test/middleware"
	"majoo-be-test/model"
	"majoo-be-test/repository"
	"majoo-be-test/transport"
	"net/http"

	"github.com/apex/log"
	"github.com/labstack/echo/v4"
)

type UserController interface {
	Login(ctx echo.Context) error
	Logout(ctx echo.Context) error
}

type userControllerImpl struct {
	config   *config.Config
	userRepo repository.UserRepo
}

func NewUserController(cfg *config.Config, userRepo repository.UserRepo) UserController {
	return &userControllerImpl{cfg, userRepo}
}

func (c *userControllerImpl) Login(ctx echo.Context) error {
	req := new(transport.UserLogin)

	if err := ctx.Bind(req); err != nil {
		log.Errorf("error: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := &model.User{
		UserName: req.Username,
		Password: req.Password,
	}

	if err := c.userRepo.Login(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Login error please check your username and password")
	}

	if err := middleware.GenerateTokensAndSetCookies(user.ID, c.config, ctx); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to generate token")
	}
	ctx.Response().Header().Set("Authorization", "<JWT>")

	return ctx.JSON(http.StatusOK, transport.BaseResponse{
		Success:  true,
		Response: "login success",
	})
}

func (c *userControllerImpl) Logout(ctx echo.Context) error {
	middleware.Logout(ctx)
	return ctx.JSON(http.StatusOK, transport.BaseResponse{
		Success:  true,
		Response: "logout success",
	})
}
