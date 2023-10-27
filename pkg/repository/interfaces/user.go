package interfaces

import (
	"sportscorner/pkg/domain"
	"sportscorner/pkg/utils/models"
)

type UserRepository interface {
	UserSignUp(user models.UserDetails) (models.UserDetailsResponse, error)
	CheckUserAvailability(email string) bool
	FindUserByEmail(user models.UserLogin) (models.UserSignInResponse, error)
	AddAdress(id int, address models.AddAddress, result bool) error
	CheckIfFirstAddress(id int) bool
	GetAddresses(id int) ([]domain.Address, error)
	GetUserDetails(id int) (models.UserDetailsResponse, error)
	
	ChangePassword(id int, password string) error
	GetPassword(id int) (string, error)
	FindIdFromPhone(phone string) (int, error)

	EditName(id int, name string) error 
	EditEmail(id int, email string) error 
	EditPhone(id int, phone string) error 

	GetCart(id int) ([]models.GetCart, error)
	RemoveFromCart(id,inv_id int) error
	UpdateQuantityAdd(id, inv_id int) error
	UpdateQuantityLess(id, inv_id int) error


	GetCartID(id int) (int, error)
	GetProductsInCart(cart_id int) ([]int, error)
	FindProductNames(inventory_id int) (string, error)
	FindCartQuantity(cart_id, inventory_id int) (int, error)
	FindPrice(inventory_id int) (float64, error)
	FindCategory(inventory_id int) (int, error)
}
