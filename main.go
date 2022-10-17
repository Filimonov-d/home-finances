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

func InputProfit(p Profit) {
	fmt.Println("Enter amount(int):")
	fmt.Scanln(p.Amount)

	fmt.Println("Enter Source(string):")
	fmt.Scanln(p.Source)

	fmt.Println("Enter Date(string):")
	fmt.Scanln(p.Date)

	fmt.Println("Ur Amount: ", p.Amount)
	fmt.Println("Ur Source: ", p.Source)
	fmt.Println("Ur Date: ", p.Date)

}

func main() {

	InputProfit(Profit{})

	/* router := gin.Default()
	router.GET("/track/listen", listenTrack)
	router.GET("/albums/listen", listenAlbums)

	router.Run("localhost:8080")
	fmt.Println("s") */
}
