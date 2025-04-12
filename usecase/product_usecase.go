package usecase

import "github.com/MarcosAndreLS/go-rest-api/model"

type ProductUsecase struct{
	//repository
}

func NewProductUseCase() ProductUsecase{
	return ProductUsecase{}
}

func (pu *ProductUsecase) GetProducts() ([]model.Product, error){
	return []model.Product{}, nil
}