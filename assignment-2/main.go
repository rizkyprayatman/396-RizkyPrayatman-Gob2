package main

import (
	"assignment-2/config"
	"assignment-2/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.DbInit()
	inDB := &controllers.InDB{DB: db}

	r := gin.Default()

	r.POST("/order", inDB.PostData)
	r.PUT("/order/:orderId", inDB.UpdateData)
	r.GET("/orders", inDB.GetAllData)
	r.DELETE("/order/:id", inDB.DeleteData)

	r.Run(":8080")
}
