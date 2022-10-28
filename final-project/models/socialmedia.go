package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	gorm.Model
	Name           string `json:"name" form:"name" valid:"required"`
	SocialMediaUrl string `json:"social_media_url" form:"social_media_url" valid:"required"`
	UserID         uint   `gorm:"foreignKey:User_Id" json:"user"`
	User           *User
}

func (sm *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(sm)

	if errCreate != nil {
		err = errCreate
		return
	}
	return
}
