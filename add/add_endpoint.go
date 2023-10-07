package add

import (
	"Fetch_API/types"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func AddHandler(c *gin.Context) {

	//we are dealing with global variables so im going to use a lock just be sure we don't mess up any data transfers
	types.PayerMutex.Lock()
	defer types.PayerMutex.Unlock()

	var RequestData types.Transaction
	err := c.BindJSON(&RequestData)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	//using a helper function to validate our points are greater than 0
	// and our payer has a name
	err = validateJSON(RequestData)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	//comparing our time string with the proper UTC RFC339 format
	//in the time library the function parse uses this 2006 specific date as a format
	// this specific date is considred "YYYY-MM-DDTHH:MM:SSZ" and we compare to make sure we have the exact
	//same format or we throw an error for improper time format
	parsed_time, err := time.Parse("2006-01-02T15:04:05Z", RequestData.Timestamp)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Incorrect Time Format")
		return
	}
	//after we validate all of our data we now have to store it
	// we are going to put our parsed time into a time.Time variable so we can sort which transactions came first
	//we will add our struct full of valid data into a slice
	RequestData.ParsedTime = parsed_time
	types.TransactionHistory = append(types.TransactionHistory, RequestData)

	//we have to store payer specific data in a map according to the assignment so we use the payer name as the key and the points at the value
	//I'm also using a global variable called TotalPointsBalance
	//TotalPointsBalance is just a quick way to check if a /spend endpoint is valid before we continue in that endpoint
	types.PayerBalances[RequestData.Payer] += RequestData.Points
	types.TotalPointsBalance += RequestData.Points

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
