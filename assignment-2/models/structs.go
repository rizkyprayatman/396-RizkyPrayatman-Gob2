package models

import "time"

type Order struct {
	OrderId      int       `json:"orderId" gorm:"primaryKey;autoIncrement:true"`
	CustomerName string    `json:"customerName"`
	OrderedAt    time.Time `json:"orderedAt" gorm:"autoCreateTime"`
	Items        []Item    `json:"items" gorm:"foreignKey:Order_ID"`
}

type Item struct {
	ItemId      int    `json:"itemId" gorm:"primaryKey;autoIncrement:true"`
	ItemCode    string `json:"itemCode" gorm:"not null;type:varchar(25)"`
	Description string `json:"description" gorm:"not null; type:text"`
	Quantity    int    `json:"quantity"`
	Order_ID    int    `json:"orderId" gorm:"not null;"`
}
