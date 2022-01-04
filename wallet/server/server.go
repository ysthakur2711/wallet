package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"wallet/api"
	"wallet/service"
	"wallet/store"
)

const (
	serverAddress = "127.0.0.1:8000"
)

// Start starts the external server
func Start() {

	go StartServer()
}

func dbConn() (db *gorm.DB) {
	/*dbDriver := "mysql"
	dbUser := "root"
	dbPass := "password"
	dbName := "walletDB"*/
	//db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/walletDB")
	db, err := gorm.Open("mysql", "root:password@tcp(127.0.0.1:3306)/walletDB?parseTime=true")
	logs.Print("db connection opened")
	if err != nil {
		panic(err.Error())
	}
	return db
}

func createRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(httprate.LimitByIP(100, 1*time.Minute))

	return r
}

func initRoutes(db *gorm.DB, r *chi.Mux) *chi.Mux {

	userRepo := store.NewUserRepo(db)
	userSvc := service.NewUserService(userRepo)
	userApi := api.NewUserResource(userSvc)

	transRepo := store.NewTransRepo(db)
	walletRepo := store.NewWalletRepo(db, transRepo)
	walletSvc := service.NewWalletService(walletRepo)
	walletApi := api.NewWalletResource(walletSvc)
	//Routes
	//public
	//userApi.RegisterRoutes(r.With(httprate.LimitByIP(10, 1*time.Minute)))
	r.Post("/users", userApi.Create)
	r.Get("/users/login", userApi.Login)

	r.Post("/wallet/addWallet", walletApi.AddWallet)
	r.Post("/wallet/pay", walletApi.Pay)
	r.Put("/wallet/credit", walletApi.Credit)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, "deepak")
	})
	return r
}

// Start starts the internal server
func StartServer() {
	log.Print("Starting server")

	db := dbConn()
	defer db.Close()

	r := createRouter()
	r = initRoutes(db, r)

	server := &http.Server{Addr: serverAddress, Handler: r}
	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	// Listen for syscall signals for process to interrupt/quit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		// Shutdown signal with grace period of 30 seconds
		shutdownCtx, _ := context.WithTimeout(serverCtx, 30*time.Second)

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()

		// Trigger graceful shutdown
		err := server.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatal("", err)
		}
		serverStopCtx()
	}()

	// Run the server
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal("failed to start server", err)
		os.Exit(1)
	}
	// Wait for server context to be stopped
	<-serverCtx.Done()
}
