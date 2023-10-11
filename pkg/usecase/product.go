package usecase

import (
	"errors"
	"sportscorner/pkg/domain"
	"sportscorner/pkg/repository/interfaces"
	"sportscorner/pkg/usecase/services"
	"sportscorner/pkg/utils/models"
)

type productUsecase struct {
	repository interfaces.ProductRepository
}

func NewProductUsecase(repo interfaces.ProductRepository) services.ProductUsecase {
	return &productUsecase{
		repository: repo,
	}
}

func (p *productUsecase) AddProduct(product domain.Product) (models.ProductResponse, error) {

	ProductResponse, err := p.repository.AddProduct(product)

	if err != nil {
		return models.ProductResponse{}, err
	}
	return ProductResponse, nil
}

func (p *productUsecase) UpdateProduct(pid int, stock int) (models.ProductResponse, error) {

	result, err := p.repository.CheckProduct(pid)
	if err != nil {
		return models.ProductResponse{}, err
	}

	if !result {
		return models.ProductResponse{}, errors.New("there is no inventory as you mentioned")
	}

	newcat, err := p.repository.UpdateProduct(pid, stock)
	if err != nil {
		return models.ProductResponse{}, err
	}

	return newcat, err
}

func (p *productUsecase) DeleteProduct(productID string) error {

	err := p.repository.DeleteProduct(productID)
	if err != nil {
		return err
	}
	return nil
}
func (p *productUsecase) ListProducts(page int, count int) ([]domain.Product, error) {

	productDetails, err := p.repository.ListProducts(page, count)
	if err != nil {
		return []domain.Product{}, err
	}

	return productDetails, nil

}

func (i *productUsecase) ShowIndividualProducts(id string) (domain.Product, error) {

	product, err := i.repository.ShowIndividualProducts(id)
	if err != nil {
		return domain.Product{}, err
	}

	return product, nil

}
