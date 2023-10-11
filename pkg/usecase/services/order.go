package services

import "sportscorner/pkg/domain"

type OrderUseCase interface {
	GetOrders(id int) ([]domain.Order, error)
	OrderItemsFromCart(userid int, addressid int, paymentid int, couponID int) error
	CancelOrder(id int) error

}
