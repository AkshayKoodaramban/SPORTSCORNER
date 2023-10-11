package services

import (
	"sportscorner/pkg/domain"
	"sportscorner/pkg/utils/models"
)

type ProductUsecase interface {
	AddProduct(product domain.Product) (models.ProductResponse, error)
	UpdateProduct(ProductID int, Stock int) (models.ProductResponse, error)
	DeleteProduct(id string) error
	ListProducts(page, count int) ([]domain.Product, error)
	ShowIndividualProducts(sku string) (domain.Product, error)
}
