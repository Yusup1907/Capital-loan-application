package usecase

import (
	"fmt"
	"net/http"
	"pinjam-modal-app/apperror"
	"pinjam-modal-app/model"
	"pinjam-modal-app/repository"
)

type ProductUsecase interface {
	CreateProduct(product *model.ProductModel) error
	GetAllProduct() ([]*model.ProductModel, error)
	GetProductById(id int) (*model.ProductModel, error)
	UpdateProduct(id int, updateProduct *model.ProductModel) error
	DeleteProduct(id int) error
}

type productUsecase struct {
	repo repository.ProductRepo
}

func (p *productUsecase) CreateProduct(product *model.ProductModel) error {
	if err := product.ValidateUpdate(); err != nil {
		return apperror.NewAppError(http.StatusBadRequest, err.Error())
	}

	productByName, err := p.repo.GetProductByName(product.ProductName)
	if err != nil {
		return fmt.Errorf("productUsecase.CreateProduct(): %w", err)
	}

	if productByName != nil {
		return apperror.AppError{
			ErrorCode:    2,
			ErrorMassage: fmt.Sprintf("Data produk dengan nama %v sudah ada", product.ProductName),
		}
	}

	return p.repo.CreateProduct(product)
}

func (p *productUsecase) GetAllProduct() ([]*model.ProductModel, error) {
	return p.repo.GetAllProduct()
}

func (p *productUsecase) GetProductById(id int) (*model.ProductModel, error) {
	return p.repo.GetProductById(id)
}

func (p *productUsecase) UpdateProduct(id int, updateProduct *model.ProductModel) error {

	if err := updateProduct.ValidateUpdate(); err != nil {
		return apperror.NewAppError(http.StatusBadRequest, err.Error())
	}

	existingProduct, err := p.repo.GetProductById(id)
	if err != nil {
		return fmt.Errorf("productUsecase.UpdateProduct(): %w", err)
	}

	if existingProduct == nil {
		return apperror.AppError{
			ErrorCode:    2,
			ErrorMassage: fmt.Sprintf("Data product dengan id %v tidak ada", id),
		}
	}

	if updateProduct.ProductName != existingProduct.ProductName {

		duplicateProduct, err := p.repo.GetProductByName(updateProduct.ProductName)
		if err != nil {
			return fmt.Errorf("productUsecase.UpdateProduct(): %w", err)
		}

		if duplicateProduct != nil {
			return apperror.AppError{
				ErrorCode:    3,
				ErrorMassage: fmt.Sprintf("Data product dengan nama product %v sudah ada", updateProduct.ProductName),
			}
		}
	}

	existingProduct.Description = updateProduct.Description
	existingProduct.Price = updateProduct.Price
	existingProduct.Stok = updateProduct.Stok
	existingProduct.CategoryProductId = updateProduct.CategoryProductId
	existingProduct.Status = updateProduct.Status

	return p.repo.UpdateProduct(id, existingProduct)
}

func (p *productUsecase) DeleteProduct(id int) error {
	existingProduct, err := p.repo.GetProductById(id)
	if err != nil {
		return fmt.Errorf("productUsecase.DeleteProduct(): %w", err)
	}

	if existingProduct == nil {
		return apperror.AppError{
			ErrorCode:    2,
			ErrorMassage: fmt.Sprintf("Data product dengan id %v tidak ada", id),
		}
	}

	return p.repo.DeleteProduct(id)
}

func NewProductUseCase(repo repository.ProductRepo) ProductUsecase {
	return &productUsecase{
		repo: repo,
	}
}