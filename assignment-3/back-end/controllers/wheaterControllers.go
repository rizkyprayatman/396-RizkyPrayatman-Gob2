package controllers

import (
	"back-end/models"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type InDB struct {
	DB *gorm.DB
}

func (idb *InDB) CreateWeather(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST")

	var (
		weather models.Weather
	)

	// Random number 1-20 (maaf enggak 1-100, agar dapat lihat status 1-15)
	water := rand.Intn(20)
	wind := rand.Intn(20)

	weather.Water = int64(water)
	weather.Wind = int64(wind)

	err := idb.DB.Model(&weather).Create(&weather).Error

	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"result": "Created data failed",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": weather,
		})
	}
	return
}

func (idb *InDB) GetWeather(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET")

	var (
		weather []models.Weather
	)

	idb.DB.Find(&weather).Take(&weather).Last(&weather)

	if len(weather) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Data not found!",
		})
	} else {
		var waterStatus, windStatus string

		getWeather := models.Weather{
			Water: weather[0].Water,
			Wind:  weather[0].Wind,
		}

		if getWeather.Water <= 5 {
			waterStatus = "Aman"
		} else if getWeather.Water >= 6 && getWeather.Water <= 8 {
			waterStatus = "Siaga"
		} else if getWeather.Water > 8 {
			waterStatus = "Bahaya"
		}

		if getWeather.Wind <= 6 {
			windStatus = "Aman"
		} else if getWeather.Wind >= 7 && getWeather.Wind <= 15 {
			windStatus = "Siaga"
		} else if getWeather.Wind > 15 {
			windStatus = "Bahaya"
		}

		c.JSON(http.StatusOK, gin.H{
			"Water":        getWeather.Water,
			"Wind":         getWeather.Wind,
			"waterMessage": waterStatus,
			"windMessage":  windStatus,
		})

	}

}

func (idb *InDB) UpdateWeather(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST")

	var (
		weather models.Weather
	)

	water := rand.Intn(20)
	wind := rand.Intn(20)

	// Random number 1-20 (maaf enggak 1-100, agar dapat lihat status 1-15)
	weather.Water = int64(water)
	weather.Wind = int64(wind)

	err := idb.DB.Model(&weather).Create(&weather).Error

	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"result": "Created data failed",
		})
	}

	var waterStatus, windStatus string

	if weather.Water <= 5 {
		waterStatus = "Aman"
	} else if weather.Water >= 6 && weather.Water <= 8 {
		waterStatus = "Siaga"
	} else if weather.Water > 8 {
		waterStatus = "Bahaya"
	}

	if weather.Wind <= 6 {
		windStatus = "Aman"
	} else if weather.Wind >= 7 && weather.Wind <= 15 {
		windStatus = "Siaga"
	} else if weather.Wind > 15 {
		windStatus = "Bahaya"
	}

	c.JSON(http.StatusOK, gin.H{
		"status":       weather,
		"waterMessage": waterStatus,
		"windMessage":  windStatus,
	})

}
