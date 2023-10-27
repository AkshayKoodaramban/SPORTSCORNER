package usecase

import (
	"fmt"
	"log"
	"sportscorner/pkg/repository/interfaces"
	"sportscorner/pkg/usecase/services"
	"sportscorner/pkg/utils/models"
	"strconv"

	"github.com/razorpay/razorpay-go"
)

type paymentUsecase struct {
	repository interfaces.PaymentRepository
}

func NewPaymentUseCase(repo interfaces.PaymentRepository) services.PaymentUseCase {
	return &paymentUsecase{
		repository: repo,
	}
}

func (p *paymentUsecase) MakePaymentRazorPay(orderID string, userID string) (models.OrderPaymentDetails, error) {

	var orderDetails models.OrderPaymentDetails
	//get orderid
	newid, err := strconv.Atoi(orderID)
	if err != nil {
		return models.OrderPaymentDetails{}, err
	}
	orderDetails.OrderID = newid

	//get userid
	newuserid, err := strconv.Atoi(userID)
	if err != nil {
		return models.OrderPaymentDetails{}, err
	}
	orderDetails.UserID = newuserid

	//get username
	username, err := p.repository.FindUsername(newuserid)
	if err != nil {
		return models.OrderPaymentDetails{}, err
	}
	orderDetails.Username = username

	//get total
	newfinal, err := p.repository.FindPrice(newid)
	if err != nil {
		return models.OrderPaymentDetails{}, err
	}

	orderDetails.FinalPrice = newfinal
	client := razorpay.NewClient("rzp_test_YFcgEYkDD4GAOh", "89xll4vuGit4Zbm4w3ogD3Cp")
	if client == nil {
		log.Fatal("empty client")
	}

	fmt.Println("final price", orderDetails.FinalPrice)

	data := map[string]interface{}{
		"amount":   int(orderDetails.FinalPrice) * 100,
		"currency": "INR",
		"receipt":  "some_receipt_id",
	}

	body, err := client.Order.Create(data, nil)
	if err != nil {
		return models.OrderPaymentDetails{}, err
	}

	razorPayOrderID := body["id"].(string)

	orderDetails.Razor_id = razorPayOrderID

	fmt.Println("order",orderDetails)
	return orderDetails, nil
}

func (p *paymentUsecase) VerifyPayment(paymentID string, razorID string, orderID string) error {

	err := p.repository.UpdatePaymentDetails(orderID, paymentID, razorID)
	if err != nil {
		return err
	}

	return nil

}
