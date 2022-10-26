package main

import (
	"back-end/config"
	"back-end/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.DbInit()
	inDB := &controllers.InDB{DB: db}

	r := gin.Default()

	r.POST("/", inDB.CreateWeather)
	r.GET("/", inDB.GetWeather)
	r.POST("/weather", inDB.UpdateWeather)

	r.Run(":8080")
}
