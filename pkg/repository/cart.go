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

    query := `
        SELECT
            p.product_name,
            li.quantity,
            (li.quantity * p.price) AS total_price
        FROM
            line_items AS li
        JOIN
            products AS p ON li.inventory_id = p.id
        WHERE
            li.user_id = ?
    `

    if err := ad.DB.Raw(query, id).Scan(&cart).Error; err != nil {
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
