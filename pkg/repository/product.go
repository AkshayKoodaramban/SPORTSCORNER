package repository

import (
	"errors"
	"fmt"
	"sportscorner/pkg/domain"
	"sportscorner/pkg/repository/interfaces"
	"sportscorner/pkg/utils/models"
	"strconv"

	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(DB *gorm.DB) interfaces.ProductRepository {
	return &ProductRepository{
		DB: DB,
	}
}

func (p *ProductRepository) AddProduct(product domain.Product) (models.ProductResponse, error) {
	var id uint
	query := `
	INSERT INTO products (category_id, product_name, size, stock, price)
    VALUES (?, ?, ?, ?, ?)
    RETURNING id
    `
	p.DB.Raw(query, product.CategoryID, product.ProductName, product.Size, product.Stock, product.Price).Scan(&id)

	var productsResponse models.ProductResponse
	productsResponse.ProductID = int(id)

	p.DB.Raw("SELECT stock FROM products WHERE product_name =?", product.ProductName).Scan(&productsResponse)

	return productsResponse, nil
}
func(p *ProductRepository)CheckProductAvilability(pname string)bool{
	var quantity int
	err:=p.DB.Raw("SELECT COUNT(product_name) from products where product_name=?",pname).Scan(&quantity).Error
	if err!=nil{
		return false
	}
	return quantity>0
}



func (p *ProductRepository) CheckProduct(pid int) (bool, error) {
	var k int
	err := p.DB.Raw("SELECT COUNT(*) FROM products WHERE id=?", pid).Scan(&k).Error
	if err != nil {
		return false, err
	}

	if k == 0 {
		return false, err
	}

	return true, err
}

func (p *ProductRepository) UpdateProduct(pid int, stock int) (models.ProductResponse, error) {

	// Check the database connection
	if p.DB == nil {
		return models.ProductResponse{}, errors.New("database connection is nil")
	}

	// Update the
	if err := p.DB.Exec("UPDATE products SET stock = stock + $1 WHERE id= $2", stock, pid).Error; err != nil {
		return models.ProductResponse{}, err
	}

	// Retrieve the update
	var newdetails models.ProductResponse
	var newstock int
	if err := p.DB.Raw("SELECT stock FROM products WHERE id=?", pid).Scan(&newstock).Error; err != nil {
		return models.ProductResponse{}, err
	}
	newdetails.ProductID = pid
	newdetails.Stock = newstock

	return newdetails, nil
}

func (p *ProductRepository) DeleteProduct(productID string) error {
	id, err := strconv.Atoi(productID)
	if err != nil {
		return errors.New("converting into integer not happened")
	}
	fmt.Println("This is the ID:", id)

	result := p.DB.Exec("DELETE FROM products WHERE id = ?", id)

	if result.RowsAffected < 1 {
		return errors.New("no records with that ID exist")
	}

	return nil
}

func (ad *ProductRepository) ListProducts(page int, count int) ([]domain.Product, error) {
	// pagination purpose -
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * count
	var productDetails []domain.Product

	if err := ad.DB.Raw("select id,category_id,product_name,size,stock,price from products limit ? offset ?", count, offset).Scan(&productDetails).Error; err != nil {
		return []domain.Product{}, err
	}

	return productDetails, nil

}

func (p *ProductRepository) CheckStock(pid int) (int, error) {
	var k int
	err := p.DB.Raw("SELECT stock FROM products WHERE id=?", pid).Scan(&k).Error
	if err != nil {
		return 0, err
	}

	return k, nil
}

func (p *ProductRepository) CheckPrice(pid int) (float64, error) {
	var k float64
	err := p.DB.Raw("SELECT price FROM products WHERE id=?", pid).Scan(&k).Error
	if err != nil {
		return 0, err
	}

	return k, nil
}

func (i *ProductRepository) ShowIndividualProducts(id string) (domain.Product, error) {
	pid, error := strconv.Atoi(id)
	if error != nil {
		return domain.Product{}, errors.New("convertion not happened")
	}
	var product domain.Product
	err := i.DB.Raw("SELECT * FROM products WHERE id=?", pid).Scan(&product).Error
	if err != nil {
		return domain.Product{}, errors.New("error retrieved record")
	}
	return product, nil
}
