package router

import (
	"majoo-be-test/controller"

	"github.com/labstack/echo/v4"
)

type MerchantRouter struct {
	merchantController controller.MerchantController
}

func NewMerchantRouter(merchantController controller.MerchantController) MerchantRouter {
	return MerchantRouter{merchantController}
}

func (r *MerchantRouter) Mount(group *echo.Group) {
	group.GET("/:merchant_id", r.merchantController.GetMerchantOmzetByID)
	group.GET("/outlet/:outlet_id", r.merchantController.GetOutletOmzetByID)
}
