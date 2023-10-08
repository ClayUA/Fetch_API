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
	// This is the amount we were requested to spend
	var SpendAmount types.SpendRequest

	// decoding json and storing it in our Struct
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

	//update our total balance by subtracting
	types.TotalPointsBalance -= SpendAmount.Points

	//This weird looking function I found in the sort library
	//It should sort our Transaction array by most recent time so we can loop through it and reduce the points
	//from the most recent transactions. This is probably not a good time complexity but for the scope of this project
	//I think should work fine
	sort.Slice(types.TransactionHistory, func(i, j int) bool {
		return types.TransactionHistory[i].ParsedTime.Before(types.TransactionHistory[j].ParsedTime)
	})

	// going to set a standard variable to our points to spend
	// I only used the SpendAmount struct to get the json information on line 25
	SpendingPoints := SpendAmount.Points

	// I like how easy it is to store payer and their balances in a map so I am using a local map to store spending in this scope
	// It should probably be persistent in a major program but it works for this project
	var SpendMap = make(map[string]int)

	// iterate through our slice of structs sorted by time
	// break if we spend the targeted amount
	for _, trans := range types.TransactionHistory {
		if SpendingPoints == 0 {
			break
		}
		// we have to see if current struct has enough points to cover the transaction
		// if we do then we assign it with the total spending points and it will only iterate once
		CurrentTransactionPoints := trans.Points
		if CurrentTransactionPoints > SpendingPoints {
			CurrentTransactionPoints = SpendingPoints
		}
		// if our Spending points are more than than the current transaction points we take them all and move to the next
		// this section subtracts the points from the current struct and our SpendingPoint value
		// we hten store our negative amounts in a map with payerID key and the points spent as the value
		trans.Points -= CurrentTransactionPoints
		SpendingPoints -= CurrentTransactionPoints
		SpendMap[trans.Payer] -= CurrentTransactionPoints
	}
	// after we spend all of our points we have to update our global map
	// this combines the spendign map values and the PayerBalances
	for PayerID, PointsSpent := range SpendMap {
		types.PayerBalances[PayerID] += PointsSpent

	}

	// This returns the amounts that were spent and the associated payer in indented JSON formate
	c.IndentedJSON(http.StatusOK, SpendMap)

}

// this just throws an error if we are trying to spend more points than we have
func ValidatePoints(obj types.SpendRequest) error {
	if obj.Points > types.TotalPointsBalance {
		return errors.New("")
	}
	return nil

}
