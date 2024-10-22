package requestmodels

import "product_app/service/model"

type ProductRequestModel struct {
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	Discount float32 `json:"discount"`
	Store    string  `json:"store"`
}

func (prm ProductRequestModel) ToModel() model.ProductCreateModel {
	return model.ProductCreateModel{
		Name:     prm.Name,
		Price:    prm.Price,
		Discount: prm.Discount,
		Store:    prm.Store,
	}
}
