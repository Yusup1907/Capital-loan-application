package usecase

import (
	"fmt"

	"pinjam-modal-app/apperror"
	"pinjam-modal-app/model"
	"pinjam-modal-app/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	GetUserById(int) (*model.UserModel, error)
	InsertUser(usr *model.UserModel) error
	UpadateUser(usr *model.UserModel) error
	GetUserByName(string) (*model.UserModel, error)
	GetAllUser() (*[]model.UserModel, error)
	DeleteUser(*model.UserModel) error
}

type userUsecaseImpl struct {
	usrRepo repository.UserRepo
}

func (usrUsecase *userUsecaseImpl) GetUserById(id int) (*model.UserModel, error) {
	return usrUsecase.usrRepo.GetUserById(id)
}

func (usrUsecase *userUsecaseImpl) GetAllUser() (*[]model.UserModel, error) {
	return usrUsecase.usrRepo.GetAllUser()
}

func (usrUsecase *userUsecaseImpl) GetUserByName(name string) (*model.UserModel, error) {
	return usrUsecase.usrRepo.GetUserByName(name)
}

func (usrUseCase *userUsecaseImpl) InsertUser(usr *model.UserModel) error {
	if usr.UserName == "" {
		return apperror.AppError{
			ErrorCode:    1,
			ErrorMessage: "Name cannot be empty",
		}
	}
	if usr.Password == "" {
		return apperror.AppError{
			ErrorCode:    1,
			ErrorMessage: "Password cannot be empty",
		}
	}

	existData, err := usrUseCase.usrRepo.GetUserByName(usr.UserName)
	if err != nil {
		return fmt.Errorf("userUsecaseImpl.InsertUser(): %w", err)
	}
	if existData != nil {
		return apperror.AppError{
			ErrorCode:    1,
			ErrorMessage: fmt.Sprintf("User data with the name %v already exists", usr.UserName),
		}
	}

	passHash, err := bcrypt.GenerateFromPassword([]byte(usr.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("userUsecaseImpl.GenerateFromPassword(): %w", err)
	}
	usr.Password = string(passHash)
	usr.RolesName = "User"
	usr.IsActive = true
	return usrUseCase.usrRepo.InsertUser(usr)
}

func (usrUseCase *userUsecaseImpl) UpadateUser(usr *model.UserModel) error {
	if usr.UserName == "" {
		return apperror.AppError{
			ErrorCode:    1,
			ErrorMessage: "Name cannot be empty",
		}
	}
	if usr.Password == "" {
		return apperror.AppError{
			ErrorCode:    1,
			ErrorMessage: "Password cannot be empty",
		}
	}

	existData, err := usrUseCase.usrRepo.GetUserById(usr.Id)
	if err != nil {
		return fmt.Errorf("userUseCaseImpl.EditUserById(): %w", err)
	}
	if existData == nil {
		return apperror.AppError{
			ErrorCode:    1,
			ErrorMessage: fmt.Sprintf("User data with the ID %v does not exist", usr.Id),
		}
	}

	existDataUsr, err := usrUseCase.usrRepo.GetUserByName(usr.UserName)
	if err != nil {
		return fmt.Errorf("userUseCaseImpl.GetUserByName(): %w", err)
	}
	if existDataUsr != nil && existDataUsr.Id != usr.Id {
		return apperror.AppError{
			ErrorCode:    1,
			ErrorMessage: fmt.Sprintf("User data with the username %v already exists", usr.UserName),
		}
	}

	passHash, err := bcrypt.GenerateFromPassword([]byte(usr.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("userUsecaseImpl.GenerateFromPassword(): %w", err)
	}
	usr.Password = string(passHash)
	usr.RolesName = "user"
	usr.IsActive = true
	return usrUseCase.usrRepo.UpadateUser(usr)
}

func (usrUsecase *userUsecaseImpl) DeleteUser(usr *model.UserModel) error {
	return usrUsecase.usrRepo.DeleteUser(usr)
}

func NewUserUseCase(usrRepo repository.UserRepo) UserUsecase {
	return &userUsecaseImpl{
		usrRepo: usrRepo,
	}
}
