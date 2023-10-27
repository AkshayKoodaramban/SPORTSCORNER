package repository

import (
	"errors"
	"fmt"
	"sportscorner/pkg/domain"
	"sportscorner/pkg/repository/interfaces"
	"sportscorner/pkg/utils/models"

	"gorm.io/gorm"
)

type UserDatabase struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &UserDatabase{DB}
}

// check whether the user is already present in the database
func (c *UserDatabase) CheckUserAvailability(email string) bool {

	var count int
	query := fmt.Sprintf("select count(*) from users where email='%s'", email)
	if err := c.DB.Raw(query).Scan(&count).Error; err != nil {
		return false
	}

	return count > 0

}

func (c *UserDatabase) UserSignUp(user models.UserDetails) (models.UserDetailsResponse, error) {

	var userDetails models.UserDetailsResponse
	err := c.DB.Raw("INSERT INTO users (name, email, password, phone) VALUES (?, ?, ?, ?) RETURNING id, name, email, phone", user.Name, user.Email, user.Password, user.Phone).Scan(&userDetails).Error

	if err != nil {
		return models.UserDetailsResponse{}, err
	}

	return userDetails, nil
}

func (c *UserDatabase) FindUserByEmail(user models.UserLogin) (models.UserSignInResponse, error) {
	var userDetails models.UserSignInResponse

	err := c.DB.Raw(`SELECT * FROM users WHERE email=? AND BLOCKED=false`, user.Email).Scan(&userDetails).Error

	if err != nil {
		return models.UserSignInResponse{}, errors.New("error checking user details")
	}
	return userDetails, nil
}

func (c *UserDatabase) AddAdress(id int, address models.AddAddress, result bool) error {
	err := c.DB.Exec(`
		INSERT INTO addresses (user_id, name, house_name, street, city, state, pin,"default")
		VALUES ($1, $2, $3, $4, $5, $6, $7,$8 )`,
		id, address.Name, address.HouseName, address.Street, address.City, address.State, address.Pin, result).Error
	if err != nil {
		return errors.New("could not add address")
	}

	return nil
}

func (c *UserDatabase) CheckIfFirstAddress(id int) bool {

	var count int
	// query := fmt.Sprintf("select count(*) from addresses where user_id='%s'", id)
	if err := c.DB.Raw("select count(*) from addresses where user_id=$1", id).Scan(&count).Error; err != nil {
		return false
	}
	// if count is greater than 0 that means the user already exist
	return count > 0

}
func (ad *UserDatabase) GetAddresses(id int) ([]domain.Address, error) {

	var addresses []domain.Address

	if err := ad.DB.Raw("select * from addresses where user_id=?", id).Scan(&addresses).Error; err != nil {
		return []domain.Address{}, errors.New("error in getting addresses")
	}

	return addresses, nil
}
func (ad *UserDatabase) GetUserDetails(id int) (models.UserDetailsResponse, error) {

	var details models.UserDetailsResponse

	if err := ad.DB.Raw("select id,name,email,phone from users where id=?", id).Scan(&details).Error; err != nil {
		return models.UserDetailsResponse{}, errors.New("could not get user details")
	}

	return details, nil

}

func (i *UserDatabase) EditName(id int, name string) error {
	err := i.DB.Exec(`update users set name=$1 where id=$2`, name, id).Error
	if err != nil {
		return err
	}
	return nil
}
func (i *UserDatabase) EditEmail(id int, email string) error {
	err := i.DB.Exec(`update users set email=$1 where id=$2`, email, id).Error
	if err != nil {
		return err
	}

	return nil
}

func (i *UserDatabase) EditPhone(id int, phone string) error {
	err := i.DB.Exec(`update users set phone=$1 where id=$2`, phone, id).Error
	if err != nil {
		return err
	}

	return nil
}
func (ad *UserDatabase) GetCartID(id int) (int, error) {

	var cart_id int

	if err := ad.DB.Raw("select id from carts where user_id=?", id).Scan(&cart_id).Error; err != nil {
		return 0, err
	}

	return cart_id, nil

}

func (ad *UserDatabase) GetProductsInCart(cart_id int) ([]int, error) {

	var cart_products []int

	if err := ad.DB.Raw("select inventory_id from line_items where cart_id=?", cart_id).Scan(&cart_products).Error; err != nil {
		return []int{}, err
	}
	fmt.Println(cart_products)
	return cart_products, nil

}

func (ad *UserDatabase) FindProductNames(inventory_id int) (string, error) {

	var product_name string
	fmt.Println("iam here 1")

	if err := ad.DB.Raw("select product_name from products where id=?", inventory_id).Scan(&product_name).Error; err != nil {
		return "", err
	}

	return product_name, nil

}

func (ad *UserDatabase) FindCartQuantity(cart_id, inventory_id int) (int, error) {

	var quantity int
	fmt.Println(cart_id, inventory_id)
	if err := ad.DB.Raw("select quantity from line_items where cart_id=$1 and inventory_id=$2", cart_id, inventory_id).Scan(&quantity).Error; err != nil {
		return 0, err
	}

	return quantity, nil

}

func (ad *UserDatabase) FindPrice(inventory_id int) (float64, error) {

	var price float64

	if err := ad.DB.Raw("select price from products where id=?", inventory_id).Scan(&price).Error; err != nil {
		return 0, err
	}

	return price, nil
}

func (ad *UserDatabase) FindCategory(inventory_id int) (int, error) {

	var category int
	fmt.Println("iam here 4")

	if err := ad.DB.Raw("select category_id from products where id=?", inventory_id).Scan(&category).Error; err != nil {
		return 0, err
	}

	return category, nil

}

func (ad *UserDatabase) GetCart(id int) ([]models.GetCart, error) {

	var cart []models.GetCart
	if err := ad.DB.Raw("select products.product_name,line_items.quantity,line_items.total_price from line_items inner join products on line_items.inventory_id=products.id where user_id=?", id).Scan(&cart).Error; err != nil {
		return []models.GetCart{}, err
	}

	return cart, nil

}
func (ad *UserDatabase) RemoveFromCart(id, inv int) error {
	if err := ad.DB.Exec("DELETE FROM line_items WHERE cart_id IN (SELECT id FROM carts WHERE user_id = $1) AND inventory_id = $2", id, inv).Error; err != nil {
		return err
	}

	return nil

}

func (ad *UserDatabase) UpdateQuantityAdd(id, inv_id int) error {
	if err := ad.DB.Exec("UPDATE line_items SET quantity = quantity + 1 WHERE cart_id IN (SELECT id FROM carts WHERE user_id = $1) AND inventory_id = $2", id, inv_id).Error; err != nil {
		return err
	}

	return nil
}

func (ad *UserDatabase) UpdateQuantityLess(id, inv_id int) error {

	// Check the current quantity
	var currentQuantity int
	// err := ad.DB.
	// 	Model(&models.LineItem{}).
	// 	Where("cart_id IN (SELECT id FROM carts WHERE user_id = ?) AND inventory_id = ?", id, inv_id).
	// 	Pluck("quantity", &currentQuantity).
	// 	Error
	// if err != nil {
	// 	return err
	// }
	if err :=ad.DB.Exec("SELECT quantity FROM line_items li JOIN carts c ON li.cart_id = c.id WHERE c.user_id = $1 AND li.inventory_id = $2",id, inv_id).Scan(&currentQuantity).Error;err!=nil{
		return err
	}

	// Check if the current quantity is zero
	if currentQuantity == 0 {
		return errors.New("Quantity is already zero, cannot decrement further")
	}

	// Perform the decrement update
	errs := ad.DB.Exec("UPDATE line_items SET quantity = quantity - 1 WHERE cart_id IN (SELECT id FROM carts WHERE user_id = ?) AND inventory_id = ?", id, inv_id).Error
	if errs != nil {
		return errs
	}

	return nil

}
func (ad *UserDatabase) ChangePassword(id int, password string) error {
	err := ad.DB.Exec("UPDATE users SET password=$1 WHERE id=$2", password, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (i *UserDatabase) GetPassword(id int) (string, error) {

	var userPassword string
	err := i.DB.Raw("select password from users where id = ?", id).Scan(&userPassword).Error
	if err != nil {
		return "", err
	}
	return userPassword, nil

}

func (ad *UserDatabase) FindIdFromPhone(phone string) (int, error) {

	var id int

	if err := ad.DB.Raw("select id from users where phone=?", phone).Scan(&id).Error; err != nil {
		return id, err
	}

	return id, nil

}
