package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type ProductUsecase struct {
	repository repository.ProductRepository
}

func NewProductUsecase(repo repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: repo,
	}
}

func (pu *ProductUsecase) Getproducts() ([]model.Product, error) {
	return pu.repository.Getproducts()
}

func (pu *ProductUsecase) CreatProduct(product model.Product) (model.Product, error) {
	productID, err := pu.repository.CreatProduct(product)
	if err != nil {
		return model.Product{}, err
	}

	product.ID = productID

	return product, err
}

func (pu *ProductUsecase) GetproductsById(id_product int) (*model.Product, error) {
	product, err := pu.repository.GetproductsById(id_product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (pu *ProductUsecase) UpdateProduct(id int, product model.Product) error {
	return pu.repository.UpdateProduct(id, product)
}

func (pu *ProductUsecase) DeleteProduct(id int) error {
	return pu.repository.DeleteProduct(id)
}
