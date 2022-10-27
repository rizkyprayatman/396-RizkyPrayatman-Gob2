package models

import (
	"errors"
	"final-project/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"not null" json:"username" form:"username" valid:"required"`
	Email    string `gorm:"not null" json:"email" form:"email" valid:"required,email"`
	Password string `gorm:"not null" json:"password"form:"password" valid:"required, minstringlength(6)"`
	Age      int    `gorm:"not null" json:"age" form:"age" valid:"required"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	if u.Age < 8 {
		return errors.New("your age must be at least 8")
	}

	var user User

	inputUsername := tx.Model(&User{}).Where("username = ?", u.Username).First(&user).Error
	inputEmail := tx.Model(&User{}).Where("email = ?", u.Email).First(&user).Error
	if inputUsername == nil {
		return errors.New("username already exist")
	} else if inputEmail == nil {
		return errors.New("email already exist")
	}

	if len(u.Password) <= 6 {
		return errors.New("password length must be at least 6 character")
	}

	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}
