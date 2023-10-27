package models

type ProductResponse struct {
	ProductID int
	Stock     int
}

type ProductUpdate struct {
	Productid int `json:"product_id"`
	Stock     int `json:"stock"`
}

type AddToCart struct {
	UserID      int `json:"user_id"`
	InventoryID int `json:"inventory_id"`
}

type Products struct {
	ID          uint    `json:"id" gorm:"unique;not null"`
	Category    string  `json:"Category"`
	ProductName string  `json:"product_name"`
	Size        string  `json:"size"`
	Stock       int     `json:"stock"`
	Price       float64 `json:"price"`
}

type AddInventories struct {
	ID          uint    `json:"id"`
	CategoryID  int     `json:"category_id"`
	ProductName string  `json:"product_name"`
	Size        string  `json:"size"`
	Stock       int     `json:"stock"`
	Price       float64 `json:"price"`
}

type SetNewName struct {
	Current string `json:"current"`
	New     string `json:"new"`
}

type Order struct {
	AddressID       int `json:"address_id"`
	PaymentMethodID int `json:"payment_id"`
}
