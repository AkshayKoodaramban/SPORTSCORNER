package services

import "sportscorner/pkg/domain"

type OrderUseCase interface {
	GetOrders(id int) ([]domain.Order, error)
	OrderItemsFromCart(userid int, addressid int, paymentid int) error
	CancelOrder(id int) error
	EditOrderStatus(status string, id int) error
	AdminOrders() (domain.AdminOrdersResponse, error)


}
