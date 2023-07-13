package usecase

import (
	"fmt"
	"pinjam-modal-app/model"
	"pinjam-modal-app/repository"
)

type GoodsUsecase interface {
	InsertGoods(*model.GoodsModel) error
	GetGoodsById(int) (*model.GoodsModel, error)
}

type goodsUsecaseImpl struct {
	goodsRepo repository.GoodsRepo
}

func (goodsUsecase *goodsUsecaseImpl) InsertGoods(goods *model.GoodsModel) error {
	customerDB, err := goodsUsecase.goodsRepo.GetCustomerById(goods.CustomerId)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	if customerDB.NIK.Valid && customerDB.NIK.String != "" &&
		customerDB.NoKK.Valid && customerDB.NoKK.String != "" &&
		customerDB.EmergencyName.Valid && customerDB.EmergencyName.String != "" &&
		customerDB.EmergencyContact.Valid && customerDB.EmergencyContact.String != "" &&
		customerDB.LastSalary.Valid && customerDB.LastSalary.Float64 != 0 {
		goods.Status = "APPROVE"
	} else {
		goods.Status = "DENIED"
		fmt.Println("Silakan lengkapi data customer")
	}

	err = goodsUsecase.goodsRepo.InsertGoods(goods)
	if err != nil {
		return fmt.Errorf("failed to insert goods: %v", err)
	}

	return nil
}

func (goodsUsecas *goodsUsecaseImpl) GetGoodsById(id int) (*model.GoodsModel, error){
	return goodsUsecas.goodsRepo.GetGoodsById(id)
}

func NewGoodsUsecase(goodsRepo repository.GoodsRepo) GoodsUsecase {
	return &goodsUsecaseImpl{
		goodsRepo: goodsRepo,
	}
}
