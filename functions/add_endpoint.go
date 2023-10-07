package functions

import (
	"Fetch_API/types"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddHandler(c *gin.Context) {

	types.PayerMutex.Lock()
	defer types.PayerMutex.Unlock()

	var RequestData types.Transaction
	err := c.BindJSON(&RequestData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = validateJSON(RequestData)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	types.TransactionHistory = append(types.TransactionHistory, RequestData)
	types.PayerBalances[RequestData.Payer] += RequestData.Points
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
