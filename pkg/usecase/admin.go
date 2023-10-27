package usecase

import (
	"errors"
	"fmt"
	"sportscorner/pkg/domain"
	helper_interface "sportscorner/pkg/helper/interface"
	"sportscorner/pkg/repository/interfaces"
	"sportscorner/pkg/usecase/services"
	"sportscorner/pkg/utils/models"

	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

type adminUsecase struct {
	adminRepository interfaces.AdminRepository
	Helper          helper_interface.Helper
}

func NewAdminUsecase(repo interfaces.AdminRepository, h helper_interface.Helper) services.AdminUseCase {
	return &adminUsecase{
		adminRepository: repo,
		Helper:          h,
	}
}

func (ad *adminUsecase) LoginHandler(adminDetails models.AdminLogin) (domain.TokenAdmin, error) {

	//geting details from the admin based on the email provided
	adminCompareDetails, err := ad.adminRepository.LoginHandler(adminDetails)
	if err != nil {
		return domain.TokenAdmin{}, errors.New("error getting details")
	}

	fmt.Println(adminCompareDetails)

	//compare password from database
	err = bcrypt.CompareHashAndPassword([]byte(adminCompareDetails.Password), []byte(adminDetails.Password))
	if err != nil {
		return domain.TokenAdmin{}, errors.New("error in comparing password with hash")
	}

	var AdminDetailsresponse models.AdminDetailsResponse
	//copy details except password to response
	err = copier.Copy(&AdminDetailsresponse, &adminCompareDetails)
	if err != nil {
		return domain.TokenAdmin{}, errors.New("some issues")
	}

	tokenString, expirationTime, err := ad.Helper.GenerateTokenAdmin(AdminDetailsresponse)
	if err != nil {
		// You can handle the error here or propagate it up, depending on your needs.
		// In this example, we're propagating the error up.
		return domain.TokenAdmin{}, err
	}

	return domain.TokenAdmin{
		Admin:  AdminDetailsresponse,
		Token:  tokenString,
		Expire: expirationTime,
	}, nil

}

func (a *adminUsecase) Blockuser(id string) error {
	user, err := a.adminRepository.GetUserByID(id)
	if err != nil {
		return err
	}
	if user.Blocked {
		return errors.New("alredy blocked")
	}

	err = a.adminRepository.UpdateBlockUserByID(user)
	if err != nil {
		return err
	}
	return nil
}

func (ad *adminUsecase) UnBlockUser(id string) error {

	user, err := ad.adminRepository.GetUserByID(id)
	if err != nil {
		return err
	}

	if user.Blocked {
		user.Blocked = false
	} else {
		return errors.New("already unblocked")
	}

	err = ad.adminRepository.UpdateUnBlockUserByID(user)
	if err != nil {
		return err
	}
	return nil
}

func (ad *adminUsecase) GetUsers(page int, count int) ([]models.UserDetailsAtAdmin, error) {

	userDetails, err := ad.adminRepository.GetUsers(page, count)
	if err != nil {
		return []models.UserDetailsAtAdmin{}, err
	}

	return userDetails, nil
}

func (i *adminUsecase) NewPaymentMethod(inv string) error {

	err := i.adminRepository.NewPaymentMethod(inv)
	if err != nil {
		return err
	}

	return nil

}