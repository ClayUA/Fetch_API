package functions

import (
	"Fetch_API/types"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddHandler(c *gin.Context) {

	var RequestData types.Transaction
	err := c.BindJSON(&RequestData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(RequestData.Payer)
	fmt.Println(RequestData.Points)
	fmt.Println(RequestData.Timestamp)

	err = validateJSON(RequestData)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
}

func validateJSON(t types.Transaction) error {
	if t.Points <= 0 {
		return fmt.Errorf("Invalid Point size. Must be greater than 0")
	}
	if t.Payer == "" {
		return fmt.Errorf("Payer name must be included")
	}
	return nil
}
