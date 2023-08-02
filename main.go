package main

import (
	"fmt"
	"strconv"
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

	basket[0] = "eagle"

	fmt.Println("Total:", basket)

	basket[0] = "0.50"

	num, err := strconv.ParseFloat(basket[0], 64)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Value: %v", num)

}
