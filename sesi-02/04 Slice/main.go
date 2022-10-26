package main

import (
	"fmt"
	"strings"
)

func main() {
	//  Slice (make function)

	var fruits = make([]string, 3)

	_ = fruits

	fmt.Printf("%#v", fruits)

	// Slice (append function)

	var buah = make([]string, 3)

	buah = append(buah, "apple", "banana", "mango")

	fmt.Printf("%#v", buah)

	// Slice (append function with ellipsis)

	var buah2 = []string{"apple", "banana", "mango"}

	var buah3 = []string{"durian", "pineapple", "startfruit"}

	buah2 = append(buah2, buah3...)

	fmt.Printf("%#v", buah2)

	// Slice (copy function)

	nn := copy(buah2, buah3)

	fmt.Println("Copied element =>", nn)

	// Slice (Slicing)
	/*
			Ada cara lain lagi agar kita dapat mendapatkan element-element dari
		sebuah slice dan kita juga bisa menentukan element dari index ke
		berapa yang ingin kita dapatkan. Caranya adalah dengan
		menggunakan slicing

		Cara penulisan dalam melakukan slicing adalah sama dengan [start:stop]. Start sama
		dengan awal index yang ingin kita akses dan stop berarti index
		akhirnya.
	*/

	var buah4 = []string{"apple", "banana", "mango", "durian", "pineapple"}

	var buah5 = buah4[1:4]
	fmt.Printf("%#v\n", buah5)

	var buah6 = buah4[0:]
	fmt.Printf("%#v\n", buah6)

	var buah7 = buah4[:3]
	fmt.Printf("%#v\n", buah7)

	var buah8 = buah4[:]
	fmt.Printf("%#v\n", buah8)

	// Slice ( Combining & Append )
	//  Kita juga dapat mengkombinasikan fungsi append dengan slicing.

	var buah9 = []string{"apple", "banana", "mango", "durian", "pineapple"}

	buah9 = append(buah9[:3], "rambutan")

	fmt.Printf("%#v\n", buah9)

	// Slice ( Backing Array )
	/*
			Setiap kita membuat suatu slice pada bahasa Go , secara otomatis Go akan membuat suatu array tersembunyi yang
		disebut dengan Backing array. Backing array akan bertugas untuk menyimpan element pada slice, bukan slice nya sendiri.
		Bahasa Go mengimplementasikan slice sebagai sebuah struktur data yang disebut dengan slice header. Slice header
		terdiri dari:
		- Alamat memori/address dari backing array.
		- Panjang dari slice yang bisa didapatkan dari fungsi len.
		- Kapasitas dari slice yang bisa didapatkan dari fungsi cap
	*/
	var buah10 = []string{"apple", "mango", "durian", "banana", "starfruit"}

	var buah11 = buah10[2:4]

	buah11[0] = "rambutan"

	fmt.Println("buah 10 => ", buah10)
	fmt.Println("buah 11 => ", buah11)

	// Slice ( Cap Array )
	/*
		Fungsi cap dapat kita gunakan untuk mengetahui kapasitas
		dari sebuah array maupun slice.
	*/
	var buah12 = []string{"apple", "mango", "durian", "banana"}

	fmt.Println("buah12 cap", cap(buah12))
	fmt.Println("buah12 len", len(buah12))

	fmt.Println(strings.Repeat("#", 20))

	var buah13 = buah12[0:3]

	fmt.Println("buah13 cap", cap(buah13))
	fmt.Println("buah13 len", len(buah13))

	fmt.Println(strings.Repeat("#", 20))

	var buah14 = buah12[1:]

	fmt.Println("buah14 cap", cap(buah14))
	fmt.Println("buah14 len", len(buah14))

	fmt.Println(strings.Repeat("#", 20))

	// Slice ( Creating a new backing )
	cars := []string{"Ford", "Honda", "Audi", "Range Rover"}
	newCars := []string{}

	newCars = append(newCars, cars[0:2]...)

	cars[0] = "Nissan"
	fmt.Println("cars:", cars)
	fmt.Println("newCars:", newCars)

}
