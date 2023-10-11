package usecase

import (
	"errors"
	"fmt"
	"sportscorner/pkg/config"
	"sportscorner/pkg/domain"
	helper_interface "sportscorner/pkg/helper/interface"
	"sportscorner/pkg/repository/interfaces"
	"sportscorner/pkg/usecase/services"
	"sportscorner/pkg/utils/models"

	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	UserRepo            interfaces.UserRepository
	Cfg                 config.Config
	otpRepository       interfaces.OtpRepository
	inventoryRepository interfaces.ProductRepository
	orderRepository     interfaces.OrderRepository
	Helper              helper_interface.Helper
}

func NewUserUseCase(repo interfaces.UserRepository, cfg config.Config, otp interfaces.OtpRepository, inv interfaces.ProductRepository, order interfaces.OrderRepository, h helper_interface.Helper) services.UserUseCase {
	return &UserUseCase{
		UserRepo:            repo,
		Cfg:                 cfg,
		otpRepository:       otp,
		inventoryRepository: inv,
		orderRepository:     order,
		Helper:              h,
	}

}

func (u *UserUseCase) UserSignUp(user models.UserDetails) (models.TokenUsers, error) {
	userExist := u.UserRepo.CheckUserAvailability(user.Email)
	if userExist {
		return models.TokenUsers{}, errors.New("user alredy exist")
	}
	if user.Password != user.ConfirmPassword {
		return models.TokenUsers{}, errors.New("passwords does not match")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return models.TokenUsers{}, errors.New("internal server error")
	}
	user.Password = string(hashedPassword)

	// add user details to the database
	userData, err := u.UserRepo.UserSignUp(user)
	if err != nil {
		return models.TokenUsers{}, err
	}

	//create JWT TOKEN string for the user
	tokenString, err := u.Helper.GenerateTokenClients(userData)
	if err != nil {
		return models.TokenUsers{}, err
	}
	var userDetils models.UserDetailsResponse
	err = copier.Copy(&userDetils, &userData)
	if err != nil {
		return models.TokenUsers{}, err
	}
	return models.TokenUsers{
		Users: userDetils,
		Token: tokenString,
	}, nil
}

func (u *UserUseCase) LoginHandler(user models.UserLogin) (models.TokenUsers, error) {
	// checking if a username exist with this email adress
	ok := u.UserRepo.CheckUserAvailability(user.Email)
	if !ok {
		return models.TokenUsers{}, errors.New("user does not exist")
	}

	// Get the user details in order to check the password, in this case ( The same function can be reused in future )
	user_details, err := u.UserRepo.FindUserByEmail(user)
	if err != nil {
		return models.TokenUsers{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user_details.Password), []byte(user.Password))
	if err != nil {
		return models.TokenUsers{}, errors.New("password incorrect")
	}

	var userDetails models.UserDetailsResponse
	err = copier.Copy(&userDetails, &user_details)
	if err != nil {
		return models.TokenUsers{}, err
	}

	tokenString, err := u.Helper.GenerateTokenClients(userDetails)
	if err != nil {
		return models.TokenUsers{}, errors.New("could not create token")
	}

	return models.TokenUsers{
		Users: userDetails,
		Token: tokenString,
	}, nil
}
func (i *UserUseCase) AddAddress(id int, address models.AddAddress) error {

	rslt := i.UserRepo.CheckIfFirstAddress(id)
	var result bool

	if !rslt {
		result = true
	} else {
		result = false
	}

	err := i.UserRepo.AddAdress(id, address, result)
	if err != nil {
		return errors.New("error in adding address")
	}

	return nil

}
func (i *UserUseCase) GetAddresses(id int) ([]domain.Address, error) {

	addresses, err := i.UserRepo.GetAddresses(id)
	if err != nil {
		return []domain.Address{}, errors.New("error in getting addresses")
	}

	return addresses, nil

}
func (i *UserUseCase) GetUserDetails(id int) (models.UserDetailsResponse, error) {

	details, err := i.UserRepo.GetUserDetails(id)
	if err != nil {
		return models.UserDetailsResponse{}, errors.New("error in getting details")
	}

	return details, nil

}

func (i *UserUseCase) EditName(id int, name string) error {

	err := i.UserRepo.EditName(id, name)
	if err != nil {
		return errors.New("could not change")
	}

	return nil
}

func (i *UserUseCase) EditEmail(id int, email string) error {

	err := i.UserRepo.EditEmail(id, email)
	if err != nil {
		return errors.New("could not change")
	}

	return nil

}

func (i *UserUseCase) EditPhone(id int, phone string) error {

	err := i.UserRepo.EditPhone(id, phone)
	if err != nil {
		return errors.New("could not change")
	}

	return nil

}

func (u *UserUseCase) GetCart(id int) ([]models.GetCart, error) {

	//find cart id
	cart_id, err := u.UserRepo.GetCartID(id)
	if err != nil {
		return []models.GetCart{}, errors.New("internal error to find cart")
	}
	//find products inide cart
	products, err := u.UserRepo.GetProductsInCart(cart_id)
	if err != nil {
		return []models.GetCart{}, errors.New("internal error find product")
	}
	//find product names
	var product_names []string
	for i := range products {
		product_name, err := u.UserRepo.FindProductNames(products[i])
		if err != nil {
			return []models.GetCart{}, errors.New("internal error find product name")
		}
		product_names = append(product_names, product_name)
	}

	//find quantity
	var quantity []int
	for i := range products {
		q, err := u.UserRepo.FindCartQuantity(cart_id, products[i])
		if err != nil {
			return []models.GetCart{}, errors.New("internal error find quantity")
		}
		fmt.Println(q)
		quantity = append(quantity, q)
	}

	var price []float64
	for i := range products {
		q, err := u.UserRepo.FindPrice(products[i])
		if err != nil {
			return []models.GetCart{}, errors.New("internal error find price")
		}
		price = append(price, q)
	}

	var categories []int
	for i := range products {
		c, err := u.UserRepo.FindCategory(products[i])
		if err != nil {
			return []models.GetCart{}, errors.New("internal error find category")
		}
		categories = append(categories, c)
	}

	var getcart []models.GetCart
	for i := range product_names {
		var get models.GetCart
		get.ProductName = product_names[i]
		get.Category_id = categories[i]
		get.Quantity = quantity[i]
		get.Total = price[i]
		get.DiscountedPrice = 0

		getcart = append(getcart, get)
	}

	// //find offers
	// var offers []int
	// for i := range categories {
	// 	c, err := u.userRepo.FindofferPercentage(categories[i])
	// 	if err != nil {
	// 		return []models.GetCart{}, errors.New("internal error")
	// 	}
	// 	offers = append(offers, c)
	// }

	// //find discounted price
	// for i := range getcart {
	// 	getcart[i].DiscountedPrice = (getcart[i].Total) - (getcart[i].Total * float64(offers[i]) / 100)
	// }

	//then return in appropriate format

	return getcart, nil

}

func (i *UserUseCase) RemoveFromCart(id int) error {

	err := i.UserRepo.RemoveFromCart(id)
	if err != nil {
		return err
	}

	return nil

}

func (i *UserUseCase) UpdateQuantityAdd(id, inv_id int) error {

	err := i.UserRepo.UpdateQuantityAdd(id, inv_id)
	if err != nil {
		return err
	}

	return nil

}

func (i *UserUseCase) UpdateQuantityLess(id, inv_id int) error {

	err := i.UserRepo.UpdateQuantityLess(id, inv_id)
	if err != nil {
		return err
	}

	return nil

}

func (i *UserUseCase) ChangePassword(id int, old string, password string, repassword string) error {

	userPassword, err := i.UserRepo.GetPassword(id)
	if err != nil {
		return errors.New("internal error")
	}

	err = i.Helper.CompareHashAndPassword(userPassword, old)
	if err != nil {
		return errors.New("password incorrect")
	}

	if password != repassword {
		return errors.New("passwords does not match")
	}

	newpassword, err := i.Helper.PasswordHashing(password)
	if err != nil {
		return errors.New("error in hashing password")
	}

	return i.UserRepo.ChangePassword(id, string(newpassword))

}

func (u *UserUseCase) ForgotPasswordSend(phone string) error {

	ok := u.otpRepository.FindUserByMobileNumber(phone)
	if !ok {
		return errors.New("the user does not exist")
	}

	u.Helper.TwilioSetup(u.Cfg.ACCOUNTSID, u.Cfg.AUTHTOKEN)
	_, err := u.Helper.TwilioSendOTP(phone, u.Cfg.SERVICESID)
	if err != nil {
		return errors.New("error ocurred while generating OTP")
	}

	return nil

}

func (u *UserUseCase) ForgotPasswordVerifyAndChange(model models.ForgotVerify) error {
	u.Helper.TwilioSetup(u.Cfg.ACCOUNTSID, u.Cfg.AUTHTOKEN)
	err := u.Helper.TwilioVerifyOTP(u.Cfg.SERVICESID, model.Otp, model.Phone)

	if err != nil {
		return errors.New("error while verifying")
	}

	id, err := u.UserRepo.FindIdFromPhone(model.Phone)
	if err != nil {
		return errors.New("cannot find user from mobile number")
	}

	newpassword, err := u.Helper.PasswordHashing(model.NewPassword)
	if err != nil {
		return errors.New("error in hashing password")
	}

	// if user is authenticated then change the password i the database
	if err := u.UserRepo.ChangePassword(id, string(newpassword)); err != nil {
		return errors.New("could not change password")
	}

	return nil
}
