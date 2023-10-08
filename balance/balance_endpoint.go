package balance

import (
	"Fetch_API/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

// This a very basic functiont hat simply return the Payer and their balances
// PayerBalances is a global variable stored in type.go that uses the payer as the key and their blaance as the value
func BalanceHandler(c *gin.Context) {
	// if we have no balances we throw an error
	if len(types.PayerBalances) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No Balances found"})
	} else {
		// if we do have balances we use an easy function to return 200OK status and return them in indented JSON format
		c.IndentedJSON(http.StatusOK, types.PayerBalances)

	}

}
