package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
		Pointer

		Pointer merupakan sebuah variable spesial yang
		digunakan untuk menyimpan alamat memori variable lainnya.

		Variabel-variabel yang memiliki reference atau alamat memori yang sama, saling berhubungan satu sama lain dan nilainya
		pasti sama. Ketika ada perubahan nilai, maka akan memberikan efek kepada variabel lain (yang alamat memorinya sama) yaitu
		nilainya juga ikut berubah.
		Cara membuat suatu variable menjadi sebuah pointer cukup mudah. Caranya adalah dengan menggunakan tanda asterisk *
		sebelum penulisan tipe datanya.
	*/

	//Memory Address
	/*
		● Kita dapat mendapatkan alamat memori dari variable
		biasa dengan menggunakan tanda ampersand &.
		● Kita juga dapat mendapatkan nilai asli dari variable
		pointer dengan cara menggunakan tanda asterisk *.
	*/
	var firstNumber int = 4
	var secondNumber *int = &firstNumber

	fmt.Println("firstNumber (value) :", firstNumber)
	fmt.Println("firstNumber (memori address) :", &firstNumber)

	fmt.Println("secondNumber (value) :", *secondNumber)
	fmt.Println("secondNumber (memori address) :", secondNumber)

	//Changing value through pointer
	/*
		Pointer digunakan untuk menyimpan alamat memori, maka
		ketika kita mengganti nilai dari sebuah pointer, maka variable
		lainnya yang mempunyai alamat memori yang sama, akan ikut
		terganti nilainya.

		Ketika nilainya berubah, tetapi perlu diingat disini bahwa alamat
		memorinya tidak berubah, hanya nilai asli yang dikandungnya
		saja. Dan juga ketika kita sebuah variable pointer ingin merubah
		nilai aslinya, maka kita juga perlu menggunakan tanda asterisk *

	*/
	var firstPerson string = "John"
	var secondPerson *string = &firstPerson

	fmt.Println("firstPerson (value) :", firstPerson)
	fmt.Println("firstPerson (memori address) :", &firstPerson)
	fmt.Println("firstPerson (value) :", *secondPerson)
	fmt.Println("firstPerson (memori address) :", secondPerson)

	fmt.Println("\n", strings.Repeat("#", 30), "\n")

	*secondPerson = "Doe"

	fmt.Println("firstPerson (value) :", firstPerson)
	fmt.Println("firstPerson (memori address) :", &firstPerson)
	fmt.Println("firstPerson (value) :", *secondPerson)
	fmt.Println("firstPerson (memori address) :", secondPerson)

	//Pointer as parameter
	/*
		Pointer juga bisa dijadikan sebagai sebuah parameter pada
		sebuah function
	*/
	var a int = 10

	fmt.Println("Before:", a)

	changeValue(&a)

	fmt.Println("After:", a)
}

func changeValue(number *int) {
	*number = 20
}
