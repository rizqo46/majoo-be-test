package server

import (
	"fmt"
	"log"
	"majoo-be-test/config"
	"majoo-be-test/router"

	customMiddleware "majoo-be-test/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	httpServer *echo.Echo
	config     *config.Config
}

func InitServer(config *config.Config) Server {
	httpServer := echo.New()
	httpServer.Use(middleware.Logger())
	return Server{httpServer, config}
}

func (s *Server) Run() {
	repository := InitRepository(s.config.DB)
	controller := InitController(s.config, repository)

	userRouter := router.NewUserRouter(controller.UserController)
	userPath := s.httpServer.Group("/user")
	userRouter.Mount(userPath)

	merchantRouter := router.NewMerchantRouter(controller.MerchantController)
	merchantPath := s.httpServer.Group("/merchant", customMiddleware.JWTmiddleware(s.config.JwtSecretKey))
	merchantRouter.Mount(merchantPath)

	// transactionRouter := router.NewTransactionRouter(controller.TransactionController)
	// transactionPath := s.httpServer.Group("/transaction", customMiddleware.JWTmiddleware(s.config.JwtSecretKey))
	// transactionRouter.Mount(transactionPath)

	if err := s.httpServer.Start(fmt.Sprintf(":%d", s.config.PORT)); err != nil {
		log.Panic(err)
	}
}
