package api

import (
	"encoding/json"
	"net/http"
	"wallet/dto"
	"wallet/service"
)

type WalletResource interface {
	AddWallet(w http.ResponseWriter, r *http.Request)
	Pay(w http.ResponseWriter, r *http.Request)
	Credit(w http.ResponseWriter, r *http.Request)
	//Get(w http.ResponseWriter, r *http.Request)
	//RegisterRoutes(r chi.Router)
}

type walletResource struct {
	walletSvc service.WalletSvc
}

func NewWalletResource(walletSvc service.WalletSvc) WalletResource {
	return &walletResource{
		walletSvc: walletSvc,
	}
}

//func (wr *walletResource) RegisterRoutes(r chi.Router) {
//	r.Get("/wallets/{walletID}", wr.Get)
//	r.Post("/wallets/pay", wr.Pay)
//}

func (wr *walletResource) AddWallet(w http.ResponseWriter, r *http.Request) {
	logrus.Println("log AddWallet in api/wallet/AddWallet ")
	var req dto.CreateWalletDto
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

	wallet, err := wr.walletSvc.AddWallet(ctx, req)
	if err != nil {
		_ = render.Render(w, r, types.ErrResponse(err))
		return
	}

	render.JSON(w, r, wallet)
}
func (wr *walletResource) Pay(w http.ResponseWriter, r *http.Request) {
	

	logrus.Println("log Pay in api/wallet/Pay ")

	var req dto.TransferMoneyDto
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

	res, err := wr.walletSvc.Pay(ctx, req)

	if err != nil {
		_ = render.Render(w, r, types.ErrResponse(err))
		return
	}

	render.JSON(w, r, res)
}

func (wr *walletResource) Credit(w http.ResponseWriter, r *http.Request) {

	logrus.Println("log Credit in api/wallet/Credit ")

	var req dto.CreditDto
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
	res, err := wr.walletSvc.Credit(ctx, req)

	if err != nil {
		_ = render.Render(w, r, types.ErrResponse(err))
		return
	}

	render.JSON(w, r, res)
}
