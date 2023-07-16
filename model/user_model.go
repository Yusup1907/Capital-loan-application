package model

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

type UserModel struct {
	Id          int        `json:"id" validate:"required"`
	UserName    string     `json:"username" validate:"required"`
	Email       string     `json:"email" validate:"required,email"`
	Password    string     `json:"password" validate:"requiredmin=12,max=12"`
	RolesName   string     `json:"roles_name" validate:"required"`
	IsActive    bool       `json:"is_active"`
	PhoneNumber string     `json:"phone_number" validate:"required,min=13,max=13"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

func (c *UserModel) Validate() error {
	validate := validator.New()
	err := validate.Struct(c)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		errMsg := ""
		for _, e := range errs {
			errMsg += fmt.Sprintf("Field %s: validation failed on tag '%s'\n", e.Field(), e.Tag())
		}
		return errors.New(errMsg)
	}

	return nil
}
