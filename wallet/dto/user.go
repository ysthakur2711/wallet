package dto

import (
	"time"
	"wallet/model"
)

type CreateUserDto struct {
	ID          int64  `json:"id" `
	Username    string `json:"username" validate:"required,alphanum"`
	Password    string `json:"password" validate:"required,min=6"`
	Email       string `json:"email" validate:"required,email"`
	Address     string `json:"address" `
	Nationality string `json:"nationality" `
	AadharNo    string `json:"aadhar_no" `
}
type UserDto struct {
	ID                int64     `json:"id"`
	Username          string    `json:"username" validate:"required,alphanum"`
	Status            string    `json:"status"`
	Email             string    `json:"email" validate:"required,email"`
	Address           string    `json:"address" `
	Nationality       string    `json:"nationality" `
	AadharNo          string    `json:"aadhar_no" `
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type LoginCredentialsDto struct {
	Username string `json:"username" validate:"required,alphanum"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoggedInUserDto struct {
	//AccessToken string  `json:"access_token"`
	User UserDto `json:"user"`
}

func NewUserDto(user model.User) UserDto {
	return UserDto{
		ID:                user.ID,
		Username:          user.Username,
		Status:            string(user.Status),
		Email:             user.Email,
		Address:           user.Address,
		Nationality:       user.Nationality,
		AadharNo:          user.AadharNo,
		PasswordChangedAt: user.PasswordChangedAt,
		CreatedAt:         user.CreatedAt,
		UpdatedAt:         user.UpdatedAt,
	}
}
