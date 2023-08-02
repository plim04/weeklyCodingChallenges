package main

import (
	"fmt"
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

func basketTotal(basket []string) int {
	basket[0] := 0.50
	basket[1] := 0.90
	basket[2] := 0.70
	basket[3] := 1.50

	for _, num := range basket {
		return
	}
}

func main() {
	basket := []string{"Banana", "Cat Food", "Bread", "Avocado"}

	fmt.Println(basketTotal(basket))
	fmt.Println("Total:", basket)
}
