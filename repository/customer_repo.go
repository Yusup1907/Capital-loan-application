package repository

import (
	"database/sql"
	"fmt"
	"pinjam-modal-app/model"
	"pinjam-modal-app/utils"
	"time"
)

type CustomerRepo interface {
	AddCustomer(*model.CustomerModel) error
	GetAllCustomer() ([]model.CustomerModel, error)
	GetCustomerById(int) (*model.CustomerModel, error)
	UpdateCustomer(*model.CustomerModel) error
	DeleteCustomer(int) error
	GetCustomerByNIK(string) (*model.CustomerModel, error)
	GetCustomerByNumber(string) (*model.CustomerModel, error)
}

type customerRepoImpl struct {
	db *sql.DB
}

func (cstRepo *customerRepoImpl) AddCustomer(cst *model.CustomerModel) error {
	qry := utils.ADD_CUSTOMER

	_, err := cstRepo.db.Exec(qry, cst.Id, cst.Full_Name, cst.Address, cst.NIK, cst.Phone_number, cst.User_Id, time.Now())
	if err != nil {
		return fmt.Errorf("error on customerRepoImpl.AddCustomer() : %w", err)
	}
	return nil
}

func (cstRepo *customerRepoImpl) GetCustomerById(id int) (*model.CustomerModel, error) {
	qry := utils.GET_CUSTOMER_BY_ID

	cst := &model.CustomerModel{}
	err := cstRepo.db.QueryRow(qry, id).Scan(&cst.Id, &cst.Full_Name, &cst.Address, &cst.NIK, &cst.Phone_number, &cst.User_Id, &cst.Created_at, &cst.Updated_at)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on customerRepoImpl.getCustomerById() : %w", err)
	}
	return cst, nil
}

func (cstRepo *customerRepoImpl) UpdateCustomer(cst *model.CustomerModel) error {
	qry := utils.UPDATE_CUSTOMER

	result, err := cstRepo.db.Exec(qry, cst.Full_Name, cst.Address, cst.NIK, cst.Phone_number, cst.User_Id, time.Now(), cst.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		return fmt.Errorf("error on customerRepoImpl.UpdateCustomer() : %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("updateService(): failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("ID tidak di temukan, ID : %d", cst.Id)
	}
	return nil
}

func (cusRepo *customerRepoImpl) GetAllCustomer() ([]model.CustomerModel, error) {
	qry := utils.GET_ALL_CUSTOMER

	cus := &model.CustomerModel{}
	rows, err := cusRepo.db.Query(qry)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on customerRepoImpl.GetAllCustomer() : %w", err)
	}

	arrCus := []model.CustomerModel{}
	for rows.Next() {
		rows.Scan(&cus.Id, &cus.Full_Name, &cus.Address, &cus.NIK, &cus.Phone_number, &cus.User_Id, &cus.Created_at, &cus.Updated_at)
		arrCus = append(arrCus, *cus)
	}
	return arrCus, nil
}

func (cstRepo *customerRepoImpl) DeleteCustomer(id int) error {

	qry := "DELETE FROM mst_customer WHERE id=$1"
	result, err := cstRepo.db.Exec(qry, id)
	if err != nil {
		return fmt.Errorf("deleteCustomer() : %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("deleteCustomer(): failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("ID tidak di temukan, ID : %d", id)
	}
	return nil
}

func (cstRepo *customerRepoImpl) GetCustomerByNIK(nik string) (*model.CustomerModel, error) {
	qry := utils.GET_CUSTOMER_BY_NIK

	cst := &model.CustomerModel{}
	err := cstRepo.db.QueryRow(qry, nik).Scan(&cst.Id, &cst.Full_Name, &cst.Address, &cst.NIK, &cst.Phone_number, &cst.User_Id, &cst.Created_at, &cst.Updated_at)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on serviceRepoImpl.GetServiceByName() : %w", err)
	}
	return cst, nil
}

func (cstRepo *customerRepoImpl) GetCustomerByNumber(phone_number string) (*model.CustomerModel, error) {
	qry := utils.GET_CUSTOMER_BY_NUMBER

	cst := &model.CustomerModel{}
	err := cstRepo.db.QueryRow(qry, phone_number).Scan(&cst.Id, &cst.Full_Name, &cst.Address, &cst.NIK, &cst.Phone_number, &cst.User_Id, &cst.Created_at, &cst.Updated_at)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on serviceRepoImpl.GetCustomerByNumber() : %w", err)
	}
	return cst, nil
}

func NewCustomerRepo(db *sql.DB) CustomerRepo {
	return &customerRepoImpl{
		db: db,
	}
}
