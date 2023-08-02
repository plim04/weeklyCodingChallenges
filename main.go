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

func main() {
	basket := []string{"Banana", "Cat Food", "Bread", "Avocado"}

	fmt.Println("Total:", basket)
}
