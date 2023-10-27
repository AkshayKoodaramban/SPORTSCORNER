package http

import (
	"sportscorner/pkg/api/handler"
	"sportscorner/pkg/api/routes"

	"github.com/gin-gonic/gin"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(userHandler *handler.UserHandler, adminHandler *handler.AdminHandler, categoryHandler *handler.CategoryHandler, productHandler *handler.ProductHandler, otpHandler *handler.OtpHandler, CartHandler *handler.CartHandler, orderHandler *handler.OrderHandler, paymentHandler *handler.PaymentHandler) *ServerHTTP {

	engine := gin.New()

	engine.LoadHTMLGlob("templates/*.html")

	// Use logger from Gin
	engine.Use(gin.Logger())

	// engine.POST("user/signup", userHandler.UserSignUp)
	// engine.POST("user/login",userHandler.LoginHandler)
	// engine.POST("admin/login",adminHandler.LoginHandler)

	routes.UserRoutes(engine.Group("/user"), userHandler, productHandler, otpHandler, CartHandler, orderHandler, paymentHandler)
	routes.AdminRoutes(engine.Group("/admin"), adminHandler, productHandler, userHandler, categoryHandler, orderHandler)

	return &ServerHTTP{engine: engine}
}

func (s *ServerHTTP) Start() error {
	err := s.engine.Run(":3000")
	if err != nil {
		return err
	}

	return nil
}
