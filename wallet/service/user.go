package service

import (
	"context"
	"github.com/sirupsen/logrus"
	"wallet/dto"
	"wallet/model"
	"wallet/pkg/local_errors"
	"wallet/store"
	"wallet/util"
)

type UserSvc interface {
	CreateUser(ctx context.Context, createUserDto dto.CreateUserDto) (dto.UserDto, error)
	LoginUser(ctx context.Context, loginCredsDto dto.LoginCredentialsDto) (dto.LoggedInUserDto, error)
}

type userService struct {
	userRepo store.UserRepo
}

func NewUserService(userRepo store.UserRepo) UserSvc {
	return &userService{
		userRepo: userRepo,
	}
}

func (u *userService) CreateUser(ctx context.Context, createUserDto dto.CreateUserDto) (dto.UserDto, error) {

	var userDto dto.UserDto

	hashedPassword, err := util.HashPassword(createUserDto.Password)

	if err != nil {
		return userDto, err
	}
	arg := store.CreateUserParams{
		Id:             createUserDto.ID,
		Username:       createUserDto.Username,
		HashedPassword: hashedPassword,
		Email:          createUserDto.Email,
		Status:         model.UserStatusACTIVE,
		Address:        createUserDto.Address,
		AadharNo:       createUserDto.AadharNo,
		Nationality:    createUserDto.Nationality,
	}
	user, err := u.userRepo.CreateUser(ctx, arg)
	if err != nil {
		return userDto, err
	}

	userDto = dto.NewUserDto(user)
	return userDto, nil
}
func (u *userService) LoginUser(ctx context.Context, loginCredentialsDto dto.LoginCredentialsDto) (dto.LoggedInUserDto, error) {
	logrus.Println("log  Login user in service/user/LoginUser ")

	var loggedInDto dto.LoggedInUserDto

	user, err := u.userRepo.GetUserByUsername(ctx, loginCredentialsDto.Username)
	if err != nil {
		return loggedInDto, err
	}

	err = util.CheckPassword(loginCredentialsDto.Password, user.HashedPassword)
	if err != nil {
		return loggedInDto, local_errors.ErrIncorrectPassword
	}

	//accessToken, err := u.tokenMaker.CreateToken(user.ID, constant.AccessTokenDuration)
	//if err != nil {
	//	return loggedInDto, err
	//}

	loggedInDto = dto.LoggedInUserDto{
		//AccessToken: accessToken,
		User: dto.NewUserDto(user),
	}
	return loggedInDto, nil
}
