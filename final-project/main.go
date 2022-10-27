package main

import (
	"final-project/config"
	"final-project/router"
)

func main() {
	config.DbInit()
	r := router.StartApp()
	r.Run(":3000")
}
