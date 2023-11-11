package handler

import (
	"fmt"
	"net/http"
	"sportscorner/pkg/usecase/services"
	"sportscorner/pkg/utils/response"

	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	usecase services.PaymentUseCase
}

func NewPaymentHandler(use services.PaymentUseCase) *PaymentHandler {
	return &PaymentHandler{
		usecase: use,
	}
}

func (p *PaymentHandler) MakePaymentRazorPay(c *gin.Context) {

	orderID := c.Query("order_id")
	userID := c.Query("user_id")
	// idValue, ok := c.Get("id")
	// if !ok {
	// 	errorRes := response.ClientResponse(http.StatusBadRequest, "check your id again", nil, nil)
	// 	c.JSON(http.StatusBadRequest, errorRes)
	// 	return
	// }
	// id, _ := idValue.(string)
	// fmt.Println("id is ", id)

	orderDetail, err := p.usecase.MakePaymentRazorPay(orderID, userID)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "could not generate order details", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	c.HTML(http.StatusOK, "razorpay.html", orderDetail)
}

func (p *PaymentHandler) VerifyPayment(c *gin.Context) {

	fmt.Println("verifying")

	orderID := c.Query("order_id")
	paymentID := c.Query("payment_id")
	razorID := c.Query("razor_id")

	fmt.Println(paymentID)
	fmt.Println(razorID)
	fmt.Println(orderID)

	err := p.usecase.VerifyPayment(paymentID, razorID, orderID)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "could not update payment details", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully updated payment details", nil, nil)
	c.JSON(http.StatusOK, successRes)

}
