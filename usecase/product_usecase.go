package usecase

import (
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
	return p.repo.CreateProduct(product)
}

func NewProductUseCase(repo repository.ProductRepo) ProductUsecase {
	return &productUsecase{
		repo: repo,
	}
}
