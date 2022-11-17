package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	ExpenceItems []ExpensesItem
	Expence      []Expense
	Profits      []Profit
	DB           *sqlx.DB
)

type Profit struct {
	Amount int    `json:"amount" db:"amount"`
	Source string `json:"source" db:"source"`
	Date   string `json:"date" db:"date"`
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
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type Expense struct {
	Date        string `json:"date" db:"date"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description"`
	Amount      int    `json:"amount" db:"amount"`
}

type Deposit struct {
	Date        string `json:"date"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Amount      int    `json:"amount"`
	ReturnDate  string `json:"returndate"`
}

func InsertDeposit(c *gin.Context) {

	fmt.Println("Inserting deposit")

	var deposit Deposit
	if err := c.ShouldBindJSON(&deposit); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	fmt.Println(deposit)
}

func InsertExpense(c *gin.Context) {

	insertSQL := "insert into expences (name,amount,date) VALUES (:name,:amount,:date)"

	fmt.Println("Inserting expense")

	var expense Expense
	if err := c.ShouldBindJSON(&expense); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	Expence = append(Expence, expense)

	res, err := DB.NamedExec(insertSQL, expense)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	fmt.Println("LastInsertId:")
	fmt.Println(res.LastInsertId())
	fmt.Println(expense)
}

func InsertExpensesItem(c *gin.Context) {
	insertSQL := "insert into expence_items (name) VALUES (:name)"

	fmt.Println("Inserting Item of expensions")

	var expensesItem ExpensesItem
	if err := c.ShouldBindJSON(&expensesItem); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	ExpenceItems = append(ExpenceItems, expensesItem)

	res, err := DB.NamedExec(insertSQL, expensesItem)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	fmt.Println("LastInsertId:")
	fmt.Println(res.LastInsertId())
	fmt.Println(ExpenceItems)
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

func InsertProfit(c *gin.Context) {

	insertSQL := "insert into profit (amount,source,date) VALUES (:amount,:source,:date)"

	fmt.Println("Inserting Profit")

	var profit Profit
	if err := c.ShouldBindJSON(&profit); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	Profits = append(Profits, profit)

	res, err := DB.NamedExec(insertSQL, profit)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	fmt.Println("LastInsertId:")
	fmt.Println(res.LastInsertId())
	fmt.Println(profit)
}

func GetExpences(c *gin.Context) {
	selectSQL := "SELECT * FROM expence_items"
	fmt.Println("Get expence items")
	if err := DB.Select(&ExpenceItems, selectSQL); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, ExpenceItems)
}

func main() {
	var err error
	dsn := "user=postgres password=postgres dbname=home_finance sslmode=disable"
	if DB, err = sqlx.Connect("postgres", dsn); err != nil {
		fmt.Println(err)
	}

	router := gin.Default()
	router.POST("/profit/insert", InsertProfit)
	router.POST("/salary/insert", InsertSalary)
	router.POST("/credit/insert", InsertCredit)
	router.POST("/expensesitem/insert", InsertExpensesItem)
	router.POST("/expense/insert", InsertExpense)
	router.POST("/deposit/insert", InsertDeposit)
	router.GET("/expenceItems", GetExpences)

	router.Run("localhost:8080")
	fmt.Println("s")
}
