package model

type employee struct {
	ID       int64  `gorm:"primary_key;AUTO_INCREMENT;not_null" json:"id"`
	Username string `json:"username;unique"`
}
