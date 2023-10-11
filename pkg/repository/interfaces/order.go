package interfaces

import (
	"sportscorner/pkg/domain"
	"sportscorner/pkg/utils/models"
)

type OrderRepository interface {
	GetOrders(id int) ([]domain.Order, error)
	GetCart(userid int) ([]models.GetCart, error)
	OrderItems(userid, addressid, paymentid int, total float64) (int, error)
	AddOrderProducts(order_id int, cart []models.GetCart) error
	CancelOrder(id int) error
}
