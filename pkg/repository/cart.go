package repository

import (
	"sportscorner/pkg/repository/interfaces"
	"sportscorner/pkg/utils/models"

	"gorm.io/gorm"
)

type cartRepository struct {
	DB *gorm.DB
}

func NewCartRepository(db *gorm.DB) interfaces.CartRepository {
	return &cartRepository{
		DB: db,
	}
}

func (ad *cartRepository) GetAddresses(id int) ([]models.Address, error) {

	var addresses []models.Address

	if err := ad.DB.Raw("SELECT * FROM addresses WHERE user_id=$1", id).Scan(&addresses).Error; err != nil {
		return []models.Address{}, err
	}

	return addresses, nil

}

func (ad *cartRepository) GetCart(id int) ([]models.GetCart, error) {
	var cart []models.GetCart

	if err := ad.DB.Raw(`
	SELECT products.product_name, line_items.quantity, (line_items.quantity * products.price) AS Total
	FROM line_items
	JOIN products ON line_items.inventory_id = products.id
	WHERE cart_id = ?
`, id).Scan(&cart).Error; err != nil {

		return []models.GetCart{}, err
	}

	return cart, nil
}

func (ad *cartRepository) GetPaymentOptions() ([]models.PaymentMethod, error) {

	var payment []models.PaymentMethod

	if err := ad.DB.Raw("SELECT * FROM payment_methods").Scan(&payment).Error; err != nil {
		return []models.PaymentMethod{}, err
	}

	return payment, nil

}

func (ad *cartRepository) GetCartId(user_id int) (int, error) {

	var id int

	if err := ad.DB.Raw("SELECT id FROM carts WHERE user_id=?", user_id).Scan(&id).Error; err != nil {
		return 0, err
	}

	return id, nil

}

func (i *cartRepository) CreateNewCart(user_id int) (int, error) {
	var id int
	// Use a parameterized query to insert the new cart
	err := i.DB.Exec("INSERT INTO carts (user_id) VALUES (?)", user_id).Error
	if err != nil {
		return 0, err
	}

	if err := i.DB.Raw("select id from carts where user_id=?", user_id).Scan(&id).Error; err != nil {
		return 0, err
	}

	return id, nil
}

func (i *cartRepository) AddLineItems(cart_id, inventory_id int) error {
	err := i.DB.Exec("INSERT INTO line_items (cart_id, inventory_id) VALUES ($1, $2)", cart_id, inventory_id).Error
	if err != nil {
		return err
	}

	return nil
}

func (i *cartRepository) CheckProductInCart(user_id, inv_id int) bool {
	var count int
	if err := i.DB.Raw("SELECT count(*) FROM line_items li JOIN carts c ON li.cart_id = c.id WHERE c.user_id = $1 AND li.inventory_id = $2", user_id, inv_id).Scan(&count).Error; err != nil {
		return false
	}
	return count > 0
}
