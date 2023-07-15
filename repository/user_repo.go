package repository

import (
	"database/sql"
	"fmt"
	"pinjam-modal-app/model"

	"pinjam-modal-app/utils"
)

type UserRepo interface {
	GetUserById(int) (*model.UserModel, error)
	GetUserByName(string) (*model.UserModel, error)
	InsertUser(*model.UserModel) error
	GetAllUser() (*[]model.UserModel, error)
	UpadateUser(usr *model.UserModel) error
	DeleteUser(*model.UserModel) error
}
type userRepoImpl struct {
	db *sql.DB
}

func (usrRepo *userRepoImpl) InsertUser(usr *model.UserModel) error {
	qry := utils.INSERT_USER

	_, err := usrRepo.db.Exec(qry, usr.Id, usr.UserName, usr.Email, usr.Password, usr.RolesName, usr.IsActive, usr.PhoneNumber, usr.CreatedAt, usr.UpdatedAt)
	if err != nil {
		return fmt.Errorf("error on userRepoImpl.InsertUser() : %w", err)
	}
	return nil
}

func (usrRepo *userRepoImpl) GetUserByName(name string) (*model.UserModel, error) {
	qry := utils.GET_USER_BY_NAME

	usr := &model.UserModel{}
	err := usrRepo.db.QueryRow(qry, name).Scan(&usr.Id, &usr.UserName, &usr.Email, &usr.Password, &usr.RolesName, &usr.IsActive, &usr.PhoneNumber, &usr.CreatedAt, &usr.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on userRepoImpl.GetuserByName() : %w", err)
	}
	return usr, nil
}

func (usrRepo *userRepoImpl) UpadateUser(usr *model.UserModel) error {
	isActiveValue := 0
	if usr.IsActive {
		isActiveValue = 1
	}

	qryid := "UPDATE mst_user SET user_name = $1, email = $2, password = $3, roles_name = $4, is_active = $5, phone_number = $6, created_at = $7, updated_at = $8 WHERE id = $9"
	_, err := usrRepo.db.Exec(qryid, usr.UserName, usr.Email, usr.Password, usr.RolesName, isActiveValue, usr.PhoneNumber, usr.CreatedAt, usr.UpdatedAt, usr.Id)
	if err != nil {
		return fmt.Errorf("error on userRepoImpl.UpdateUser(): %w", err)
	}

	return nil
}

func (usrRepo *userRepoImpl) GetUserById(id int) (*model.UserModel, error) {
	qry := utils.GET_USER_BY_ID

	usr := &model.UserModel{}
	err := usrRepo.db.QueryRow(qry, id).Scan(&usr.Id, &usr.UserName, &usr.Email, &usr.Password, &usr.RolesName, &usr.IsActive, &usr.PhoneNumber, &usr.CreatedAt, &usr.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Data pengguna dengan ID %d tidak ditemukan", id)
		}
		return nil, fmt.Errorf("error on userRepoImpl.GetUserById(): %w", err)
	}
	return usr, nil
}

func (usrRepo *userRepoImpl) GetAllUser() (*[]model.UserModel, error) {
	qry := utils.GET_ALL_USER

	rows, err := usrRepo.db.Query(qry)
	if err != nil {
		return nil, fmt.Errorf("error on UserRepoImpl.GetAllUser(): %w", err)
	}
	defer rows.Close()

	users := make([]model.UserModel, 0)
	for rows.Next() {
		usr := model.UserModel{}
		err := rows.Scan(&usr.Id, &usr.UserName, &usr.Email, &usr.Password, &usr.RolesName, &usr.IsActive, &usr.PhoneNumber, &usr.CreatedAt, &usr.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("error on UserRepoImpl.GetAllUser(): %w", err)
		}
		users = append(users, usr)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error on UserRepoImpl.GetAllUser(): %w", err)
	}

	return &users, nil
}

func (usrRepo *userRepoImpl) DeleteUser(usr *model.UserModel) error {
	qry := utils.DELETE_USER

	_, err := usrRepo.db.Exec(qry, usr.Id)
	if err != nil {
		return fmt.Errorf("error on CategoryLoanRepoImpl.DeleteCategoryLoan(): %w", err)
	}
	return nil
}

func NewUserRepo(db *sql.DB) UserRepo {
	return &userRepoImpl{
		db: db,
	}
}
