package model

import "time"

type UserModel struct {
	Id          int        `json:"id"`
	UserName    string     `json:"username"`
	Email       string     `json:"email"`
	Password    string     `json:"password"`
	RolesName   string     `json:"roles_name"`
	IsActive    bool       `json:"is_active"`
	PhoneNumber string     `json:"phone_number"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}
