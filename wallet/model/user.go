package model

import "time"

type UserStatus string

const (
	UserStatusACTIVE  UserStatus = "ACTIVE"
	UserStatusBLOCKED UserStatus = "BLOCKED"
)

type User struct {
	ID                int64      `gorm:"primary_key;AUTO_INCREMENT;not_null" json:"id"`
	Username          string     `json:"username;unique"`
	HashedPassword    string     `json:"hashed_password"`
	Status            UserStatus `json:"status"`
	Email             string     `json:"email"`
	Address           string     `json:"address"`
	Nationality       string     `json:"nationality"`
	AadharNo          string     `json:"aadhar_no"`
	PasswordChangedAt time.Time  `json:"password_changed_at"`
	CreatedAt         time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt         time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

////// util me password.go me jakr
//// HashPassword return bcrypt hash of the password
//func HashPassword(password string) (string, error) {
//	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
//	if err != nil {
//		return "", fmt.Errorf("faied to hash password: %w", err)
//	}
//	return string(hashedPassword), nil
//}
//
//// CheckPassword check if the provided password is correct or not
//func CheckPassword(password string, hashedPassword string) error {
//	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
//}
