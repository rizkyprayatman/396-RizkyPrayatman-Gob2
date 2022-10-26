package controllers

import (
	"assignment-2/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type InDB struct {
	DB *gorm.DB
}

func (idb *InDB) PostData(c *gin.Context) {
	var (
		order  models.Order
		result gin.H
	)

	c.BindJSON(&order)
	idb.DB.Create(&order)
	result = gin.H{
		"result": order,
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) GetAllData(c *gin.Context) {
	var (
		orders []models.Order
		result gin.H
	)

	err := idb.DB.Model(&models.Order{}).Preload("Items").Find(&orders).Error
	_ = err

	result = gin.H{
		"result": orders,
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) UpdateData(c *gin.Context) {
	var (
		order  models.Order
		result gin.H
	)

	id, _ := strconv.Atoi(c.Param("orderId"))
	order.OrderId = id
	c.BindJSON(&order)

	idb.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&order)
	idb.DB.Model(&order).Updates(&order)

	result = gin.H{
		"Success": "data update successfully",
		"Result":  order,
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) DeleteData(c *gin.Context) {
	var (
		order  models.Order
		result gin.H
	)

	id := c.Param("id")
	err := idb.DB.First(&order, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}

	err = idb.DB.Select("Items").Delete(&order).Error
	if err != nil {
		result = gin.H{
			"result": "delete failed",
		}
	} else {
		result = gin.H{
			"result": "data deleted successfully",
		}
	}

	c.JSON(http.StatusOK, result)
}
