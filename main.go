package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Profit struct {
	Amount int    `json:"amount"`
	Source string `json:"source"`
	Date   string `json:"date"`
}

func InsertProfit(c *gin.Context) {

	fmt.Println("Inserting Profit")

	var profit Profit
	if err := c.ShouldBindJSON(&profit); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	fmt.Println(profit)
}

func main() {

	router := gin.Default()
	router.POST("/profit/insert", InsertProfit)

	router.Run("localhost:8080")
	fmt.Println("s")
}
