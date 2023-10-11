package usecase

import (
	"sportscorner/pkg/domain"
	"sportscorner/pkg/repository/interfaces"
	"sportscorner/pkg/usecase/services"
)

type orderUseCase struct {
	orderRepository interfaces.OrderRepository
	userUseCase     services.UserUseCase
}

func NewOrderUseCase(repo interfaces.OrderRepository, userUseCase services.UserUseCase) services.OrderUseCase {
	return &orderUseCase{
		orderRepository: repo,
		userUseCase:     userUseCase,
	}
}

func (i *orderUseCase) GetOrders(id int) ([]domain.Order, error) {

	orders, err := i.orderRepository.GetOrders(id)
	if err != nil {
		return []domain.Order{}, err
	}

	return orders, nil

}

func (i *orderUseCase) OrderItemsFromCart(userid int, addressid int, paymentid int, couponID int) error {

	cart, err := i.userUseCase.GetCart(userid)
	if err != nil {
		return err
	}

	var total float64
	for _, v := range cart {
		total = total + v.Total
	}

	// //finding discount if any
	// DiscountRate := i.couponRepository.FindCouponDiscount(couponID)

	// totalDiscount := (total * float64(DiscountRate)) / 100

	// total = total - totalDiscount

	order_id, err := i.orderRepository.OrderItems(userid, addressid, paymentid, total)
	if err != nil {
		return err
	}

	if err := i.orderRepository.AddOrderProducts(order_id, cart); err != nil {
		return err
	}

	return nil

}

func (i *orderUseCase) CancelOrder(id int) error {

	err := i.orderRepository.CancelOrder(id)
	if err != nil {
		return err
	}
	return nil

}
