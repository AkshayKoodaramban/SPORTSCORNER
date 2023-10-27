package handler

import (
	"net/http"
	"sportscorner/pkg/usecase/services"
	"sportscorner/pkg/utils/models"
	"sportscorner/pkg/utils/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderUseCase services.OrderUseCase
}

func NewOrderHandler(useCase services.OrderUseCase) *OrderHandler {
	return &OrderHandler{
		orderUseCase: useCase,
	}
}

func (i *OrderHandler) GetOrders(c *gin.Context) {
	idValue, ok := c.Get("id")
	
	if !ok {
		errorRes := response.ClientResponse(http.StatusBadRequest, "check your id again", nil, nil)
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	id, _ := idValue.(int)
	orders, err := i.orderUseCase.GetOrders(id)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not retrieve records", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully got all records", orders, nil)
	c.JSON(http.StatusOK, successRes)
}

func (i *OrderHandler) OrderItemsFromCart(c *gin.Context) {
	// idstring := c.Query("coupon-id")
	// if idstring == "" {
	// 	idstring = "0"
	// }
	// couponId, err := strconv.Atoi(idstring)
	// if err != nil {
	// 	errorRes := response.ClientResponse(http.StatusBadRequest, "coupon id trouble", nil, err.Error())
	// 	c.JSON(http.StatusBadRequest, errorRes)
	// 	return
	// }
	idValue, ok := c.Get("id")
	if !ok {
		errorRes := response.ClientResponse(http.StatusBadRequest, "check your id again", nil, nil)
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	id, _ := idValue.(int)


	var order models.Order

	if err := c.BindJSON(&order); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	if err := i.orderUseCase.OrderItemsFromCart(id, order.AddressID, order.PaymentMethodID); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not make the order", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully made the order", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

func (i *OrderHandler) EditOrderStatus(c *gin.Context) {

	status := c.Query("status")
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "coonversion to integer not possible", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	if err := i.orderUseCase.EditOrderStatus(status, id); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully edited the order status", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

func (i *OrderHandler) AdminOrders(c *gin.Context) {

	orders, err := i.orderUseCase.AdminOrders()
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not retrieve records", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully got all records", orders, nil)
	c.JSON(http.StatusOK, successRes)
}
