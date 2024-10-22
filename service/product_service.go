package service

import (
	"errors"
	"product_app/domain"
	"product_app/persistence"
	"product_app/service/model"
)

type IProductService interface {
	Add(productCreate model.ProductCreateModel) error
	GetById(productId int64) (domain.Product, error)
	GetAllProducts() []domain.Product
}

type ProductService struct {
	pr persistence.IProductRepository
}

func NemProductService(pr persistence.IProductRepository) IProductService {
	return &ProductService{
		pr: pr,
	}
}

func (ps *ProductService) Add(productCreate model.ProductCreateModel) error {
	err := validateProductCreate(productCreate)
	if err != nil {
		return err
	}

	return ps.pr.AddProduct(domain.Product{
		Name:     productCreate.Name,
		Price:    productCreate.Price,
		Discount: productCreate.Discount,
		Store:    productCreate.Store,
	})
}
func (ps *ProductService) GetById(productId int64) (domain.Product, error) {
	return ps.pr.GetById(productId)
}
func (ps *ProductService) GetAllProducts() []domain.Product {
	return ps.pr.GetAllProducts()
}

func validateProductCreate(productCreate model.ProductCreateModel) error {
	if productCreate.Discount > 70 {
		return errors.New("validasyon hatası")
	}
	if productCreate.Discount < 0 {
		return errors.New("adam ol lan")
	}
	return nil
}
