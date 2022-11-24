package main

import (
	"context"
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
	Deposits     []Deposit
	Credits      []Credit
	Salaries     []Salary
	iDates       []iDate
	DB           *sqlx.DB
)

type iDate struct {
	Date string `json:"date" db:"date"`
	Sum  string `json:"sum" db:"sum"`
}

type Profit struct {
	Amount int    `json:"amount" db:"amount"`
	Source string `json:"source" db:"source"`
	Date   string `json:"date" db:"date"`
}

type Salary struct {
	Amount int    `json:"amount" db:"amount"`
	Date   string `json:"date" db:"date"`
}

type Credit struct {
	Amount      int    `json:"amount" db:"amount"`
	Date        string `json:"date" db:"date"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	ReturnDate  string `json:"returndate" db:"returndate"`
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

type CurrentMoney struct {
	Amount int `json:"amount" db:"amount"`
}

type Deposit struct {
	Date        string `json:"date" db:"date"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Amount      int    `json:"amount" db:"amount"`
	ReturnDate  string `json:"returndate" db:"returndate"`
}

func InsertDeposit(c *gin.Context) {

	insertSQL := "insert into deposit (name,description,amount,returndate,date) VALUES (:name,:description,:amount,:returndate,:date)"

	fmt.Println("Inserting deposit")

	var deposit Deposit
	if err := c.ShouldBindJSON(&deposit); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	Deposits = append(Deposits, deposit)

	res, err := DB.NamedExec(insertSQL, deposit)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	fmt.Println("LastInsertId:")
	fmt.Println(res.LastInsertId())
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

	insertSQL := "insert into credit (amount, date, name, description, returndate) VALUES (:amount,:date, :name, :description, :returndate)"

	fmt.Println("Inserting Credit")

	var credit Credit
	if err := c.ShouldBindJSON(&credit); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	Credits = append(Credits, credit)

	res, err := DB.NamedExec(insertSQL, credit)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	fmt.Println("LastInsertId:")
	fmt.Println(res.LastInsertId())
	fmt.Println(credit)
}

func InsertSalary(c *gin.Context) {

	insertSQL := "insert into salary (amount,date) VALUES (:amount,:date)"

	fmt.Println("Inserting Salary")

	var salary Salary
	if err := c.ShouldBindJSON(&salary); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	Salaries = append(Salaries, salary)

	res, err := DB.NamedExec(insertSQL, salary)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	fmt.Println("LastInsertId:")
	fmt.Println(res.LastInsertId())
	fmt.Println(salary)
}

func InsertProfit(c *gin.Context) {

	insertSQL := "insert into profit (amount,source,date) VALUES (:amount,:source,:date)"
	insertSQLl := "insert into money (amount,date) VALUES (:amount,:date)"

	var profit Profit

	ctx := context.Background()
	tx, err := DB.BeginTx(ctx, nil)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	_, err = tx.NamedExecContext(ctx, "insert into profit (amount,source,date) VALUES (:amount,:source,:date)")
	if err != nil {
		tx.Rollback()
		return
	}

	_, err = tx.NamedExecContext(ctx, "insert into money (amount,date) VALUES (:amount,:date)")
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	/*insertSQL := "insert into profit (amount,source,date) VALUES (:amount,:source,:date)"

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
	}*/

	fmt.Println("LastInsertId:")
	/*fmt.Println(res.LastInsertId())
	fmt.Println(profit)*/
	fmt.Println(tx)
}

func InsertDate(c *gin.Context) {
	var date iDate

	if err := c.ShouldBindJSON(&date); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	fmt.Println(date)

}

func GetMoney(c *gin.Context) {

	selectSQL := "select sum(m.amount)  from money m where m.date <= $1"

	fmt.Println("Get Money")

	if err := DB.Select(&iDates, selectSQL, "12.12.2022"); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	fmt.Println()

	c.JSON(http.StatusOK, ExpenceItems)
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
	router.POST("/date/insert", InsertDate)
	router.GET("/expenceItems", GetExpences)
	router.GET("/money", GetMoney)

	router.Run("localhost:8080")
	fmt.Println("s")
}
