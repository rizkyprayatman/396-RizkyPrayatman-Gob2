package main

import (
	"fmt"
	"strings"
)

func main() {
	//Declare closure in variable
	/*
		Jika suatu variable mengandung sebuah closure, maka variable
		tersebut memiliki sifat seperti closure yang dikandungnya. Cara
		untuk memanggil closure tersebut, kita perlu memanggil variable
		yang mengandungnya.
	*/
	var evenNumbers = func(numbers ...int) []int {
		var result []int

		for _, v := range numbers {
			if v%2 == 0 {
				result = append(result, v)
			}
		}

		return result
	}

	var numbers = []int{4, 93, 77, 10, 52, 22, 34}

	fmt.Println(evenNumbers(numbers...))

	//IIFE
	/*
		IIFE atau singkatan dari immediately-invoked function expression
		merupakan sebuah closure yang dapat langsung tereksekusi
		ketika pertama kali dideklarasikan
	*/
	var isPalindrome = func(str string) bool {
		var temp string = ""

		for i := len(str) - 1; i >= 0; i-- {
			temp += string(byte(str[i]))
		}

		return temp == str
	}("katak")

	fmt.Println(isPalindrome)

	//Closure as a return value
	/*
		Closure juga bisa dijadikan sebagai nilai kembalian dari suatu
		function
	*/
	var studentLists = []string{"Airell", "Nanda", "Mailo", "Schannel", "Marco"}

	var find = findStudent(studentLists)

	fmt.Println(find("airell"))

	//Callback
	/*
		Callback adalah sebuah closure yang dijadikan sebagai parameter
		pada sebuah function
	*/
	var numbers2 = []int{2, 5, 8, 10, 3, 99, 23}

	var find2 = findOddNumbers(numbers2, func(number int) bool {
		return number%2 != 0
	})

	fmt.Println("Total odd numbers:", find2)
}

// Closure as a return value
func findStudent(students []string) func(string) string {
	return func(s string) string {
		var student string
		var position int

		for i, v := range students {
			if strings.ToLower(v) == strings.ToLower(s) {
				student = v
				position = i
				break
			}
		}
		if student == "" {
			return fmt.Sprintf("%s doesn't exist!!!", s)
		}
		return fmt.Sprintf("We found %s at position %d", s, position+1)
	}
}

// Callback
func findOddNumbers(numbers []int, callback func(int) bool) int {
	var totalOddNumbers int
	for _, v := range numbers {
		if callback(v) {
			totalOddNumbers++
		}
	}
	return totalOddNumbers
}
