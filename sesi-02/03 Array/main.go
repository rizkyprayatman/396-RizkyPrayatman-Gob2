package main

import (
	"fmt"
)

func main() {
	var numbers [4]int
	numbers = [4]int{1, 2, 3, 4}

	var strings = [3]string{"Airell", "Nanda", "Mailo"}

	fmt.Printf("%#v\n", numbers)
	fmt.Printf("%#v\n", strings)

	// Array(Modify Element Through Index)

	var fruits = [3]string{"apel", "pisang", "mangga"}
	fruits[0] = "apple"
	fruits[1] = "banana"
	fruits[2] = "mango"

	fmt.Printf("%#v\n", fruits)

	// Array(Loop through elements)

	var food = [3]string{"Ayam Bakar", "Ayam Goreng", "Ayam Pop"}
	// Cara Pertama
	for i, v := range food {
		fmt.Printf("index: %d, value: %s\n", i, v)
	}

	// fmt.Println(strings.Repeat("#", 25))

	// Cara Kedua

	for i := 0; i < len(food); i++ {
		fmt.Printf("Index: %d, value: %s\n", i, food[i])
	}

	// Array(Multidimensional Array)

	balances := [2][3]int{{5, 6, 7}, {8, 9, 10}}

	for _, arr := range balances {
		for _, value := range arr {
			fmt.Printf("%d ", value)
		}
		fmt.Println()
	}
}
