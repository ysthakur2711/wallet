package api

import (
	"encoding/json"
	"github.com/go-chi/render"
	"github.com/go-playground/validator"
	"github.com/sirupsen/logrus"
	"github.com/ysthakur2711/wallet/dto"
	types "github.com/ysthakur2711/wallet/pkg/local_errors"
	"github.com/ysthakur2711/wallet/service"
	"net/http"
)

type UserResource interface {
	Create(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	//RegisterRoutes(r chi.Router)
}

type userResource struct {
	userSvc service.UserSvc
}

func NewUserResource(userSvc service.UserSvc) UserResource {
	return &userResource{
		userSvc: userSvc,
	}
}

//func (u *userResource) RegisterRoutes(r chi.Router) {
//	r.Post("/users", u.Create)
//	r.Post("/users/login", u.Login)
//}

func (u *userResource) Create(w http.ResponseWriter, r *http.Request) {

	logrus.Println("log create user in api/user/Create ")
	var req dto.CreateUserDto
	ctx := r.Context()

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		_ = render.Render(w, r, types.ErrBadRequest(err))
		return
	}
	defer r.Body.Close()

	if err := validator.New().Struct(req); err != nil {
		_ = render.Render(w, r, types.ErrBadRequest(err))
		return
	}
	user, err := u.userSvc.CreateUser(ctx, req)
	if err != nil {
		_ = render.Render(w, r, types.ErrResponse(err))
		return
	}
	render.JSON(w, r, user)
}

func (u *userResource) Login(w http.ResponseWriter, r *http.Request) {

	logrus.Println("log Login user in api/user/Login ")
	var req dto.LoginCredentialsDto
	ctx := r.Context()

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		_ = render.Render(w, r, types.ErrBadRequest(err))
		return
	}
	defer r.Body.Close()

	if err := validator.New().Struct(req); err != nil {
		_ = render.Render(w, r, types.ErrBadRequest(err))
		return
	}

	loggedInUser, err := u.userSvc.LoginUser(ctx, req)
	if err != nil {
		_ = render.Render(w, r, types.ErrResponse(err))
		return
	}

	render.JSON(w, r, loggedInUser)
}
