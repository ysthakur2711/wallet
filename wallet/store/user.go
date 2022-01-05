package store

import (
	"context"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"

	"time"
	"wallet/model"
	"wallet/pkg/local_errors"
)

type UserRepo interface {
	CreateUser(ctx context.Context, arg CreateUserParams) (model.User, error)
	GetUserByUsername(ctx context.Context, username string) (model.User, error)
	//UpdateUserStatus(ctx context.Context, arg UpdateUserStatusParams) (model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepo(client *gorm.DB) UserRepo {
	return &userRepository{
		db: client,
	}
}

type CreateUserParams struct {
	Id             int64            `gorm:"primary_key;AUTO_INCREMENT;not_null" json:"id"`
	Username       string           `json:"username"`
	HashedPassword string           `json:"hashed_password"`
	Status         model.UserStatus `json:"status"`
	Email          string           `json:"email"`
	Address        string           `json:"address"`
	Nationality    string           `json:"nationality"`
	AadharNo       string           `json:"aadhar_no"`
}

func (q *userRepository) CreateUser(ctx context.Context, arg CreateUserParams) (model.User, error) {

	logrus.Println("log create user in store/user/CreateUser ")

	var u model.User
	res := q.db.Where("username = ?", arg.Username).Take(&u)
	// SELECT * FROM users WHERE username = "jinzhu";

	// check error ErrRecordNotFound
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		u = model.User{
			ID:                arg.Id,
			Username:          arg.Username,
			HashedPassword:    arg.HashedPassword,
			Status:            arg.Status,
			Email:             arg.Email,
			Address:           arg.Address,
			Nationality:       arg.Nationality,
			AadharNo:          arg.AadharNo,
			PasswordChangedAt: time.Now(),
		}
		res = q.db.Create(&u) // pass pointer of data to Create
	} else {
		logrus.Println("username already exist !! ")
		return u, local_errors.ErrUsernameAlreadyTaken
	}
	return u, nil
}
func (q *userRepository) GetUserByUsername(ctx context.Context, username string) (model.User, error) {
	logrus.Println("log  Login user in store/user/")

	var u model.User
	res := q.db.Where("username = ?", username).Take(&u)
	// SELECT * FROM users WHERE username = "jinzhu";

	// check error ErrRecordNotFound
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		logrus.Println("username not found !! ")
		return u, fmt.Errorf("wrong username")
	}
	return u, nil
}
