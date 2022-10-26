package main

import "fmt"

func main() {
	// Loopings (ﬁrst way of looping)

	// for i := 0; i < 3; i++ {
	// 	fmt.Println("Angka", i)
	// }

	//Loopings (second way of looping)

	// var i = 0

	// for i < 3 {
	// 	fmt.Println("Angka", i)
	// 	i++
	// }

	// Loopings (third way of looping)

	// var i = 0

	// for {
	// 	fmt.Println("Angka", i)

	// 	i++
	// 	if i == 3 {
	// 		break
	// 	}
	// }

	// Loopings (break and continue keywords)

	// for i := 1; i <= 10; i++ {
	// 	if i%2 == 1 {
	// 		continue
	// 	}

	// 	if i > 8 {
	// 		break
	// 	}

	// 	fmt.Println("Angka", i)
	// }

	// Loopings (Nested Looping)

	// for i := 0; i < 5; i++ {
	// 	for j := i; j < 5; j++ {
	// 		fmt.Print(j, " ")
	// 	}

	// 	fmt.Println()
	// }

	// Loopings (Label)

outerloop:
	for i := 0; i < 3; i++ {
		fmt.Println("Perulangan ke - ", i+1)
		for j := 0; j < 3; j++ {
			if i == 2 {
				break outerloop
			}
			fmt.Print(j, " ")
		}
		fmt.Print(("\n"))
	}
}
