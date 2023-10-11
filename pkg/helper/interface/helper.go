package interfaces

import (
	"sportscorner/pkg/utils/models"
)

type Helper interface {
	GenerateTokenAdmin(admin models.AdminDetailsResponse) (string, int64, error)
	TwilioSetup(username string, password string)
	TwilioSendOTP(phone string, serviceID string) (string, error)
	TwilioVerifyOTP(serviceID string, code string, phone string) error
	GenerateTokenClients(user models.UserDetailsResponse) (string, error)
	PasswordHashing(string) (string, error)
	CompareHashAndPassword(a string, b string) error
}
