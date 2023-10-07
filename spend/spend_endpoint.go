package spend

import (
	"Fetch_API/types"
	"errors"
	"net/http"
	"sort"
	"sync"

	"github.com/gin-gonic/gin"
)

func SpendHandler(c *gin.Context) {
	var Mutex sync.Mutex
	//using a lock when operating on data
	Mutex.Lock()
	defer Mutex.Unlock()

	// creating a single variable struct to read in point intergers from json to int
	var SpendAmount types.SpendRequest

	// decoding json and applying it to our Struct
	// if the format is not correct it will return a bad request error 400
	err := c.BindJSON(&SpendAmount)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	//using our TotalPointsBalance to see if we have enough points stored to complete the transaction
	flag := ValidatePoints(SpendAmount)
	if flag != nil {
		c.JSON(http.StatusBadRequest, "User does not have enough points for Reqeust")
		return
	}
	//we have validated that we can have enough points to redeem so we will update our total balance
	types.TotalPointsBalance -= SpendAmount.Points

	//This weird looking function I found in the sort library
	//It should sort our Transaction array by most recent time so we can loop through it and reduce the points
	//from the most recent transactions. This is probably not a good time complexity but for the scope of this project
	//I think should work fine
	sort.Slice(types.TransactionHistory, func(i, j int) bool {
		return types.TransactionHistory[i].ParsedTime.Before(types.TransactionHistory[j].ParsedTime)
	})
	SpendingPoints := SpendAmount.Points
	var SpendMap = make(map[string]int)
	for _, trans := range types.TransactionHistory {
		if SpendingPoints == 0 {
			break
		}

		CurrentTransactionPoints := trans.Points
		if CurrentTransactionPoints > SpendingPoints {
			CurrentTransactionPoints = SpendingPoints
		}

		trans.Points -= CurrentTransactionPoints
		SpendingPoints -= CurrentTransactionPoints
		SpendMap[trans.Payer] -= CurrentTransactionPoints
	}
	for PayerID, PointsSpent := range SpendMap {
		types.PayerBalances[PayerID] += PointsSpent

	}
	c.IndentedJSON(http.StatusOK, SpendMap)

}

func ValidatePoints(obj types.SpendRequest) error {
	if obj.Points > types.TotalPointsBalance {
		return errors.New("")
	}
	return nil

}
