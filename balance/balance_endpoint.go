package balance

import (
	"Fetch_API/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BalanceHandler(c *gin.Context) {
	if len(types.PayerBalances) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No Balances found"})
	} else {
		c.IndentedJSON(http.StatusOK, types.PayerBalances)

	}

}
