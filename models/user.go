package models

import (
	"time"

	"Final-Project-BDS-Sanbercode-Golang-Batch-30/utils/token"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type (
	// Role struct {
	// 	ID        uint      `json:"id" gorm:"primary_key"`
	// 	Name      string    `json:"name"`
	// 	CreatedAt time.Time `json:"created_at"`
	// 	UpdatedAt time.Time `json:"updated_at"`
	// }

	User struct {
		ID       uint   `json:"id" gorm:"primary_key"`
		Name     string `json:"name"`
		Email    string `json:"email" gorm:"unique"`
		Password string `json:"password"`
		Avatar   string `json:"avatar"`
		// RoleID    uint      `json:"role_id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		// Role      Company   `json:"-"`
	}
)

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(email string, password string, db *gorm.DB) (string, error) {

	var err error

	u := User{}

	err = db.Model(User{}).Where("email = ?", email).Take(&u).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := token.GenerateToken(u.ID)

	if err != nil {
		return "", err
	}

	return token, nil

}

func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	//turn password into hash
	hashedPassword, errPassword := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if errPassword != nil {
		return &User{}, errPassword
	}
	u.Password = string(hashedPassword)

	var err error = db.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}
