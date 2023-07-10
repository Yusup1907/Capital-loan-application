package usecase

import (
	"fmt"
	"pinjam-modal-app/apperror"
	"pinjam-modal-app/model"
	"pinjam-modal-app/repository"
)

type ProductUsecase interface {
	CreateProduct(product *model.ProductModel) error
}

type productUsecase struct {
	repo repository.ProductRepo
}

func (p *productUsecase) CreateProduct(product *model.ProductModel) error {
	productByName, err := p.repo.GetProductByName(product.ProductName)
	if err != nil {
		return fmt.Errorf("productUsecase.CreateProduct() : %w", err)
	}

	if productByName != nil {
		return apperror.AppError{
			ErrorCode:    1,
			ErrorMassage: fmt.Sprintf("data product dengan nama product %v sudah ada", product.ProductName),
		}
	}
	return p.repo.CreateProduct(product)
}

func NewProductUseCase(repo repository.ProductRepo) ProductUsecase {
	return &productUsecase{
		repo: repo,
	}
}
