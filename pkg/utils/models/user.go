package models

type UserDetails struct {
	Name            string `json:"name"`
	Email           string `json:"email" validate:"email"`
	Phone           string `json:"phone"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmpassword"`
}

type Address struct {
	Id        uint   `json:"id" gorm:"unique;not null"`
	UserID    uint   `json:"user_id"`
	Name      string `json:"name" validate:"required"`
	HouseName string `json:"house_name" validate:"required"`
	Street    string `json:"street" validate:"required"`
	City      string `json:"city" validate:"required"`
	State     string `json:"state" validate:"required"`
	Pin       string `json:"pin" validate:"required"`
}

type UserDetailsResponse struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type TokenUsers struct {
	Users UserDetailsResponse
	Token string
}

type UserLogin struct {
	Email    string `json:"email" validate:"email"`
	Password string `json:"password"`
}

type UserSignInResponse struct {
	Id       uint   `json:"id"`
	UserID   uint   `json:"user_id"`
	Name     string `json:"name"`
	Email    string `json:"email" validate:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type AddAddress struct {
	Name      string `json:"name" validate:"required"`
	HouseName string `json:"house_name" validate:"required"`
	Street    string `json:"street" validate:"required"`
	City      string `json:"city" validate:"required"`
	State     string `json:"state" validate:"required"`
	Pin       string `json:"pin" validate:"required"`
}

type ChangePassword struct {
	Oldpassword string `json:"old_password"`
	Password    string `json:"password"`
	Repassword  string `json:"re_password"`
}

type EditName struct {
	Name string `json:"name"`
}

type EditEmail struct {
	Email string `json:"email"`
}

type EditPhone struct {
	Phone string `json:"phone"`
}

type GetCart struct {
	ProductName string  `json:"product_name"`
	Category_id int     `json:"category_id"`
	Quantity    int     `json:"quantity"`
	Total       float64 `json:"total_price"`
	// DiscountedPrice float64 `json:"discounted_price"`
}

type CheckOut struct {
	Addresses      []Address
	Products       []GetCart
	PaymentMethods []PaymentMethod
	TotalPrice     float64
}

type ForgotPasswordSend struct {
	Phone string `json:"phone"`
}

type ForgotVerify struct {
	Phone       string `json:"phone"`
	Otp         string `json:"otp"`
	NewPassword string `json:"newpassword"`
}
