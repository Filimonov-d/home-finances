package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Profit struct {
	Amount int    `json:"amount"` //Так дописывать в каждой структуре
	Source string `json:"source"`
	Date   string `json:"date"`
}

type Salary struct {
	Amount int    `json:"amount"`
	Date   string `json:"date"`
}

type Credit struct {
	Amount      int    `json:"amount"`
	Date        string `json:"date"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ReturnDate  string `json:"returndate"`
}

type ExpensesItem struct {
	Name string `json:"name"`
}

type Expense struct {
}

func InsertExpensesItem(c *gin.Context) {

	fmt.Println("Inserting Item of expensions")

	var expensesitem ExpensesItem
	if err := c.ShouldBindJSON(&expensesitem); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	fmt.Println(expensesitem)
}

func InsertCredit(c *gin.Context) {

	fmt.Println("Inserting Credit")

	var credit Credit
	if err := c.ShouldBindJSON(&credit); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	fmt.Println(credit)
}

func InsertSalary(c *gin.Context) {

	fmt.Println("Inserting Salary")

	var salary Salary
	if err := c.ShouldBindJSON(&salary); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	fmt.Println(salary)
}

func InsertProfit(c *gin.Context) { //с жин контекст это обязательно в каждой функции

	fmt.Println("Inserting Profit")

	var profit Profit                                 // Это переменная которая содержит структуру
	if err := c.ShouldBindJSON(&profit); err != nil { //Это выводит данные
		c.String(http.StatusBadRequest, err.Error())
	}

	fmt.Println(profit)
}

func main() {

	router := gin.Default()
	router.POST("/profit/insert", InsertProfit) // Это постит, путь, и что
	router.POST("/salary/insert", InsertSalary)
	router.POST("/credit/insert", InsertCredit)
	router.POST("/expensesitem/insert", InsertExpensesItem)

	router.Run("localhost:8080")
	fmt.Println("s")
}
