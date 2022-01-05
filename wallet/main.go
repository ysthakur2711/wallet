package wallet

import (
	logs "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"wallet/server"
)

//func dbConn() (db *gorm.DB) {
//	/*dbDriver := "mysql"
//	dbUser := "root"
//	dbPass := "password"
//	dbName := "walletDB"*/
//	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/walletDB")
//	if err != nil {
//		panic(err.Error())
//	}
//	return db
//}

//func insert() {
//
//	var u model.employee
//	res := q.db.Where("username = ?", "yogesh").Take(&u)
//
//	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
//		u = model.employee{
//			ID:       arg.Id,
//			Username: arg.Username,
//		}
//		res = q.db.Create(&u) // pass pointer of data to Create
//	} else {
//		logrus.Println("username already exist !! ")
//		return u, local_errors.ErrUsernameAlreadyTaken
//	}
//	return u, nil
//}

func main() {

	// SetFormatter sets the standard logger formatter.
	logs.SetFormatter(&logs.TextFormatter{})
	logs.Println("starting wallet service")

	//db := dbConn()
	//defer db.Close()

	server.Start()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-interrupt
}
