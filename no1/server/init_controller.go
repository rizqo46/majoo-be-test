package server

import (
	"majoo-be-test/config"
	"majoo-be-test/controller"
)

type Controller struct {
	UserController     controller.UserController
	MerchantController controller.MerchantController
}

func InitController(cfg *config.Config, repository Repository) Controller {
	return Controller{
		UserController:     controller.NewUserController(cfg, repository.UserRepo),
		MerchantController: controller.NewMerchantController(cfg, repository.MerchantRepo),
	}
}
