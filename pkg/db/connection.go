package db

import (
	"fmt"
	"sportscorner/pkg/config"
	"sportscorner/pkg/domain"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	db.AutoMigrate(&domain.Users{
		ID:           0,
		Name:         "",
		Email:        "",
		Password:     "",
		Phone:        "",
		Blocked:      false,
		IsAdmin:      false,
	})
	db.AutoMigrate(&domain.Admin{})
	db.AutoMigrate(&domain.Product{})
	db.AutoMigrate(&domain.Category{})
	db.AutoMigrate(domain.Address{})
	db.AutoMigrate(domain.LineItems{})
	db.AutoMigrate(domain.Cart{})
	db.AutoMigrate(domain.Order{
		Model:           gorm.Model{},
		UserID:          0,
		Users:           domain.Users{},
		AddressID:       0,
		Address:         domain.Address{},
		PaymentMethodID: 0,
		PaymentMethod:   domain.PaymentMethod{},
		CouponUsed:      "",
		FinalPrice:      0,
		OrderStatus:     "",
		PaymentStatus:   "",
	})
	db.AutoMigrate(domain.OrderItem{})
	db.AutoMigrate(domain.PaymentMethod{})



	CheckAndCreateAdmin(db)

	return db, dbErr
}

func CheckAndCreateAdmin(db *gorm.DB) {
	var count int64
	db.Model(&domain.Admin{}).Count(&count)
	if count == 0 {
		password := "456123009"
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return
		}

		admin := domain.Admin{
			ID:       1,
			Name:     "sportscorner",
			Username: "sportscorner@gmail.com",
			Password: string(hashedPassword),
		}
		db.Create(&admin)
	}
}
