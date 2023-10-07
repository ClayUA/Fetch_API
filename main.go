package main

import (
	"Fetch_API/functions"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/add", functions.AddHandler)
	router.GET("/balance", functions.BalanceHandler)

	router.Run(":8000")

}
