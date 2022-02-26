package controller

import (
	"majoo-be-test/config"
	"majoo-be-test/middleware"
	"majoo-be-test/repository"
	"net/http"
	"strconv"

	jwtv3 "github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type MerchantController interface {
	GetMerchantOmzetByID(ctx echo.Context) error
	GetOutletOmzetByID(ctx echo.Context) error
}

type merchantControllerImpl struct {
	config       *config.Config
	merchantRepo repository.MerchantRepo
}

func NewMerchantController(cfg *config.Config, merchantRepo repository.MerchantRepo) MerchantController {
	return &merchantControllerImpl{cfg, merchantRepo}
}

func (c *merchantControllerImpl) GetMerchantOmzetByID(ctx echo.Context) error {
	userID := getUserID(ctx)
	merchantID, _ := strconv.Atoi(ctx.Param("merchant_id"))

	ok, err := c.merchantRepo.IsUserAuthToMerchant(userID, merchantID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "cannot get data")
	}

	if !ok {
		return echo.NewHTTPError(http.StatusBadRequest, "you are not allowed to see the data")
	}

	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))
	page, _ := strconv.Atoi(ctx.QueryParam("page"))

	res, err := c.merchantRepo.GetMerchantOmzet(merchantID, limit, page)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "cannot get data")
	}
	return ctx.JSON(http.StatusOK, res)
}

func (c *merchantControllerImpl) GetOutletOmzetByID(ctx echo.Context) error {
	userID := getUserID(ctx)
	outletID, _ := strconv.Atoi(ctx.Param("outlet_id"))

	ok, err := c.merchantRepo.IsUserAuthToOutlet(userID, outletID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "cannot get data")
	}

	if !ok {
		return echo.NewHTTPError(http.StatusBadRequest, "you are not allowed to see the data")
	}

	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))
	page, _ := strconv.Atoi(ctx.QueryParam("page"))

	res, err := c.merchantRepo.GetOutletOmzet(outletID, limit, page)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "cannot get data")
	}
	return ctx.JSON(http.StatusOK, res)
}

func getUserID(c echo.Context) int {
	u := c.Get("user").(*jwtv3.Token)
	return u.Claims.(*middleware.Claims).UserID
}
