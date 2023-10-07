package main

import (
	"Fetch_API/functions"
	"Fetch_API/types"

	"github.com/gin-gonic/gin"
)

var TransactionHistory []types.Transaction

func main() {
	router := gin.Default()

	router.POST("/add", functions.AddHandler)

	router.Run(":8000")

}
