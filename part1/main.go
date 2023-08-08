package main

import (
	"fmt"
	"math"
)

//PART 1
//Create a function that returns the total (an int) for the all items in the basket
//Note that the cost of each item is:
//Banana 50p
//Cat Food 90p
//Bread 70p
//Avocado £1.50

//the total should be printed to the console

// func main() {
// 	basket := []string{"Banana", "Cat Food", "Bread", "Avocado"}

// 	basketTotal(basket)
// 	fmt.Println("Total:", basket)
// }

func main() {
	basket := []string{"Banana", "Cat Food", "Bread", "Avocado"}

	totalPrice := basketItemsTotal(basket)
	fmt.Println("Total price of items:", math.Round(totalPrice*100)/100, "GBP")

}

func basketItemsTotal(basket []string) float64 {
	basketItemPrices := map[string]float64{
		"Banana":   0.50,
		"Cat Food": 0.90,
		"Bread":    0.70,
		"Avocado":  1.50,
	}

	totalPrice := 0.0

	for _, item := range basket {
		if price, found := basketItemPrices[item]; found {
			totalPrice += price
		}
	}
	return totalPrice
}
