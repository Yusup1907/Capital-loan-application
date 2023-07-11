package usecase

import (
	"fmt"
	"pinjam-modal-app/apperror"
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
	cpDB, err := cpUsecase.cpRepo.GetCategoryProductByName(cp.CategoryProductName)
	if err != nil {
		return fmt.Errorf("serviceUsecaseImpl.InsertService() : %w", err)
	}

	if cpDB != nil {
		return apperror.AppError{
			ErrorCode:    1,
			ErrorMassage: fmt.Sprintf("data category product dengan nama %v sudah ada", cp.CategoryProductName),
		}
	}
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