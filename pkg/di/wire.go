package di

// //go:build wireinject
// // +build wireinject

import (
	http "sportscorner/pkg/api"
	"sportscorner/pkg/api/handler"
	"sportscorner/pkg/config"
	"sportscorner/pkg/db"
	"sportscorner/pkg/helper"
	"sportscorner/pkg/repository"
	"sportscorner/pkg/usecase"

	"github.com/google/wire"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(
		db.ConnectDatabase,
		helper.NewHelper,
		repository.NewAdminRepository,
		usecase.NewAdminUsecase,
		handler.NewAdminHandler,

		repository.NewCategoryRepository,
		usecase.NewCategoryUseCase,
		handler.NewCategoryHandler,

		repository.NewOtpRepository,
		usecase.NewOtpUseCase,
		handler.NewOtpHandler,

		repository.NewUserRepository,
		usecase.NewUserUseCase,
		handler.NewUserHandler,

		repository.NewProductRepository,
		usecase.NewProductUsecase,
		handler.NewProductHandler,

		repository.NewCartRepository,
		usecase.NewCartUseCase,
		handler.NewCartHandler,

		repository.NewOrderRepository,
		usecase.NewOrderUseCase,
		handler.NewOrderHandler,

		repository.NewPaymentRepository,
		usecase.NewPaymentUseCase,
		handler.NewPaymentHandler,

		http.NewServerHTTP,
	)

	return &http.ServerHTTP{}, nil
}
