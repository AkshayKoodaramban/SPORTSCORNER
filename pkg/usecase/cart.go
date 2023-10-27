package usecase

import (
	"errors"
	"sportscorner/pkg/repository/interfaces"
	"sportscorner/pkg/usecase/services"
	"sportscorner/pkg/utils/models"
)

type cartUseCase struct {
	repo              interfaces.CartRepository
	productRepository interfaces.ProductRepository
}

func NewCartUseCase(repo interfaces.CartRepository, productRepo interfaces.ProductRepository) services.CartUseCase {
	return &cartUseCase{
		repo:              repo,
		productRepository: productRepo,
	}
}

func (i *cartUseCase) AddToCart(user_id, inventory_id int) error {
	//check if the desired product has quantity available
	stock, err := i.productRepository.CheckStock(inventory_id)
	if err != nil {
		return err
	}
	//if available then call userRepository
	if stock <= 0 {
		return errors.New("out of stock")
	}

	//find user cart id
	cart_id, err := i.repo.GetCartId(user_id)
	if err != nil {
		return errors.New("some error in geting user cart")
	}
	//if user has no existing cart create new cart
	if cart_id == 0 {
		cart_id, err = i.repo.CreateNewCart(user_id)
		if err != nil {
			return errors.New("cannot create cart for user")
		}
	}

	count := i.repo.CheckProductInCart(user_id, inventory_id)
	if count {
		return errors.New("product alredy in cart")
	}

	//add product to line items
	if err := i.repo.AddLineItems(cart_id, inventory_id); err != nil {
		return errors.New("error in adding products")
	}

	return nil
}

func (i *cartUseCase) CheckOut(id int) (models.CheckOut, error) {

	address, err := i.repo.GetAddresses(id)
	if err != nil {
		return models.CheckOut{}, err
	}

	payment, err := i.repo.GetPaymentOptions()
	if err != nil {
		return models.CheckOut{}, err
	}

	cartID, err := i.repo.GetCartId(id)
	if err != nil {
		return models.CheckOut{}, err
	}

	products, err := i.repo.GetCart(cartID)
	if err != nil {
		return models.CheckOut{}, err
	}

	var price float64
	for _, v := range products {
		price = price + v.Total
	}

	var checkout models.CheckOut

	checkout.Addresses = address
	checkout.Products = products
	checkout.PaymentMethods = payment
	checkout.TotalPrice = price

	return checkout, err
}
