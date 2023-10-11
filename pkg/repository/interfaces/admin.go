package interfaces

import (
	"sportscorner/pkg/domain"
	"sportscorner/pkg/utils/models"
)

type AdminRepository interface {
	LoginHandler(adminDetails models.AdminLogin) (domain.Admin, error)
	GetUserByID(id string) (domain.Users, error)
	GetUsers(page int, count int) ([]models.UserDetailsAtAdmin, error)
	UpdateBlockUserByID(user domain.Users) error
	UpdateUnBlockUserByID(user domain.Users) error
}
