package main

import "fmt"

func main() {
	fmt.Println("Constants & Operators")
	// Constant

	const pi float32 = 3.14

	/*
		Constant (const) atau Konstanta merupakan
		jenis variable pada bahasa Go yang nilainya tidak dapat diubah.
	*/

	// Operators

	/*
		terdapat 3 jenis operator yang perlu kita ketahui
		yaitu operator aritmatika, operator logika danoperator perbandingan.
	*/

	var value = 2 + 2*3

	fmt.Println(value)

	// Operator perbandingan/relasional
	var firstCondition bool = 2 < 3
	var secondCondition bool = "joey" == "joey"
	var thirdCondition bool = 10 != 2.3
	var fourthCondition bool = 11 <= 11

	fmt.Println("first condition :", firstCondition)
	fmt.Println("second condition :", secondCondition)
	fmt.Println("third condition :", thirdCondition)
	fmt.Println("fourth condition :", fourthCondition)

	// Operator logika

	var right = true
	var wrong = false

	var wrongAndRight = wrong && right
	fmt.Printf("wrong && right \t(%t) \n", wrongAndRight)

	var wrongOrRight = wrong || right
	fmt.Printf("wrong || right \t(%t) \n", wrongOrRight)

	var wrongReverse = !wrong
	fmt.Printf("!wrong \t\t(%t) \n", wrongReverse)
}
