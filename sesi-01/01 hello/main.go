package main

import "fmt"

func main() {
	// Menghasilkan Pesan Hello World
	fmt.Println("Hello World")

	/*
		Perbedaan Println dan Print
		Println akan mengeksekusi dan baris baru
		Print mengeksekusi tidak pada baris baru
	*/

	fmt.Println("Airell Jordan")
	fmt.Println("Airell", "Jordan")

	fmt.Print("Airell Jordan\n")
	fmt.Print("Airell", " Jordan\n")
	fmt.Print("Airell", " ", "Jordan\n")
}
