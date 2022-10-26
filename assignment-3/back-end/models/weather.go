package models

import "gorm.io/gorm"

type Weather struct {
	gorm.Model
	Water int64 `gorm:"column:water;not null;type:int" json:"water"`
	Wind  int64 `gorm:"column:wind;not null;type:int" json:"wind"`
}
