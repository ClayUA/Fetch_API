package functions

import (
	"Fetch_API/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BalanceHandler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, types.PayerBalances)

}
