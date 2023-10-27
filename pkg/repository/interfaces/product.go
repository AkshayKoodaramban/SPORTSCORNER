package interfaces

import (
	"sportscorner/pkg/domain"
	"sportscorner/pkg/utils/models"
)

type ProductRepository interface {
	AddProduct(product domain.Product) (models.ProductResponse, error)
	CheckProduct(pid int) (bool, error)
	UpdateProduct(pid int, stock int) (models.ProductResponse, error)
	DeleteProduct(id string) error
	ListProducts(page int, count int) ([]domain.Product, error)
	CheckStock(product_id int) (int, error)
	CheckPrice(product_id int) (float64, error)
	ShowIndividualProducts(id string) (domain.Product, error)
	CheckProductAvilability(pname string)bool
}
