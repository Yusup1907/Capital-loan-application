package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"pinjam-modal-app/model"
	"time"

	"pinjam-modal-app/utils"
)

type UserRepo interface {
	CreateUser(user *model.UserModel) error
	GetUserByUsername(string) (*model.UserModel, error)
	GetUserByEmail(email string) (*model.UserModel, error)
	GetUserByUsernameOrEmail(identifier string) (*model.UserModel, error)
	GetUserById(int) (*model.UserModel, error)
	GetAllUser() (*[]model.UserModel, error)
	UpadateUser(usr *model.UserModel) error
	DeleteUser(*model.UserModel) error
}
type userRepoImpl struct {
	db *sql.DB
}

var (
	// ErrUserNotFound digunakan untuk menandakan pengguna tidak ditemukan dalam database.
	ErrUserNotFound = errors.New("User not found")
)

func (r *userRepoImpl) CreateUser(user *model.UserModel) error {
	insertStatement := "INSERT INTO mst_user (user_name, email, password, roles_name, is_active, phone_number, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err := r.db.Exec(insertStatement, user.UserName, user.Email, user.Password, user.RolesName, user.IsActive, user.PhoneNumber, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepoImpl) GetUserByUsername(username string) (*model.UserModel, error) {
	selectStatment := "SELECT id, user_name, email, password, roles_name, is_active, phone_number, created_at, updated_at FROM mst_user WHERE user_name = $1"

	row := r.db.QueryRow(selectStatment, username)

	user := &model.UserModel{}
	err := row.Scan(&user.Id, &user.UserName, &user.Email, &user.Password, &user.RolesName, &user.IsActive, &user.PhoneNumber, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

func (r *userRepoImpl) GetUserByEmail(email string) (*model.UserModel, error) {
	selectStatment := "SELECT id, user_name, email, password, roles_name, is_active, phone_number, created_at, updated_at FROM mst_user WHERE email = $1"

	row := r.db.QueryRow(selectStatment, email)

	user := &model.UserModel{}
	err := row.Scan(&user.Id, &user.UserName, &user.Email, &user.Password, &user.RolesName, &user.IsActive, &user.PhoneNumber, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

func (r *userRepoImpl) GetUserByUsernameOrEmail(identifier string) (*model.UserModel, error) {
	selectStatement := "SELECT id, user_name, email, password, roles_name, is_active, phone_number, created_at, updated_at FROM mst_user WHERE user_name = $1 OR email = $2"

	row := r.db.QueryRow(selectStatement, identifier, identifier)

	user := &model.UserModel{}
	err := row.Scan(&user.Id, &user.UserName, &user.Email, &user.Password, &user.RolesName, &user.IsActive, &user.PhoneNumber, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("GetUserByUsernameOrEmail() : %w", err)
	}

	return user, nil
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
