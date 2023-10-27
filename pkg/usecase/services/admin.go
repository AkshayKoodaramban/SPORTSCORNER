package services

import (
	"sportscorner/pkg/domain"
	"sportscorner/pkg/utils/models"
)

type AdminUseCase interface {
	LoginHandler(adminDetails models.AdminLogin) (domain.TokenAdmin, error)
	Blockuser(id string) error
	UnBlockUser(id string) error
	GetUsers(page int, count int) ([]models.UserDetailsAtAdmin, error)
	NewPaymentMethod(string) error
}
