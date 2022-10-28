package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Message string `json:"message" form:"message" valid:"required"`
	UserID  uint   `gorm:"foreignKey:User_Id" json:"user_id" form:"user_id" required:"true"`
	PhotoID uint   `gorm:"foreignKey:Photo_Id" json:"photo_id" form:"photo_id" required:"true"`
	User    *User
	Photo   *Photo
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(c)

	if errCreate != nil {
		err = errCreate
		return
	}
	return
}
