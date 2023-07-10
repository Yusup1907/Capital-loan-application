package usecase

import (
	"pinjam-modal-app/model"
	"pinjam-modal-app/repository"
)

type CategoryProductUsecase interface {
	InsertCategoryProduct(*model.CategoryProductModel) error
	GetCategoryProductById(int) (*model.CategoryProductModel, error)
}

type categoryProductUsecaseImpl struct {
	cpRepo repository.CategoryProductRepo
}

func (cpUsecase *categoryProductUsecaseImpl) InsertCategoryProduct(cp *model.CategoryProductModel) error{
	return cpUsecase.cpRepo.InsertCategoryProduct(cp)
}

func (cpUsecase *categoryProductUsecaseImpl) GetCategoryProductById(id int)(*model.CategoryProductModel, error){
	return cpUsecase.cpRepo.GetCategoryProductById(id)
}

func NewCategoryProductUsecase(cpRepo repository.CategoryProductRepo) CategoryProductUsecase {
	return &categoryProductUsecaseImpl{
		cpRepo: cpRepo,
	}
}