package main

import (
	"fmt"
	//"github.com/gin-gonic/gin"
)

type Profit struct {
	Amount int
	Source string
	Date   string
}

type Income struct {
	Amount int
}

var PointerProfit = Profit{}
var PointerIncome = Income{}

func InputAmount(p *Profit) {

	fmt.Println("\nEnter Amount(integer):")
	if _, err := fmt.Scanln(&p.Amount); err != nil {
		fmt.Println("Amount error: ", err)
		InputAmount(&PointerProfit)
	}
}

func InputSource(p *Profit) {

	fmt.Println("\nEnter Source(string):")
	if _, err := fmt.Scanln(&p.Source); err != nil {
		fmt.Println("Source error: ", err)
		InputSource(&PointerProfit)
	}
}

func InputDate(p *Profit) {

	fmt.Println("\nEnter Date(string):")
	if _, err := fmt.Scanln(&p.Date); err != nil {
		fmt.Println("Date error: ", err)
		InputDate(&PointerProfit)
	}
}
func InputData(p *Profit) {

	InputAmount(&PointerProfit)
	defer fmt.Println("Profit Amount: ", p.Amount)

	InputSource(&PointerProfit)
	defer fmt.Println("Profit Source: ", p.Source)

	InputDate(&PointerProfit)
	defer fmt.Println("\nProfit Date: ", p.Date)
	defer fmt.Println("\nUr Profit: ")

}

func InputIncome(i *Income) {

	fmt.Println("Enter Amount(integer):")
	if _, err := fmt.Scanln(&i.Amount); err != nil {
		fmt.Println("Amount error: ", err)
		InputIncome(&PointerIncome)
	}
	defer fmt.Println("\nMonth income amount: ", i.Amount)

}

func main() {

	InputIncome(&PointerIncome)

	InputData(&PointerProfit)

	/* router := gin.Default()
	router.GET("/track/listen", listenTrack)
	router.GET("/albums/listen", listenAlbums)

	router.Run("localhost:8080")
	fmt.Println("s") */
}
