package seed

import (
	"github.com/jinzhu/gorm"
	"wallet/model"
)

var users = []model.User{
	model.User{
		Username:       "Yogesh",
		Email:          "Yogesh@gmail.com",
		HashedPassword: "password",
	},
	model.User{
		Username:       "Ys",
		Email:          "Ysthakur@gmail.com",
		HashedPassword: "password",
	},
}

func Load(db *gorm.DB) {

}
