package main

import (
	"Fetch_API/add"
	"Fetch_API/balance"
	"Fetch_API/spend"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.POST("/add", add.AddHandler)
	router.GET("/balance", balance.BalanceHandler)
	router.POST("/spend", spend.SpendHandler)

	router.Run(":8000")

}
