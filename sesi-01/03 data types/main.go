package main

import "fmt"

func main() {
	fmt.Println("Data Types")

	/*
		Tipe data pada bahasa Go terbagi menjadi 4 kategori
		dengan detail seperti berikut ini:
		1.Basic Type: number, string, boolean.
		2.Aggregate Type: array dan struct.
		3.Reference Type: slice, pointer, map, function, channel
		4.Interface Type: interface
	*/

	// Number (interger) https://www.tutorialspoint.com/go/go_data_types.htm

	var first = 12
	var second = 89

	fmt.Printf("tipe data first : %T\n", first)
	fmt.Printf("bilangan second : %T\n", second)

	// uint8, int8

	var data1 uint8 = 22
	var data2 int8 = -10

	fmt.Printf("tipe data data1 : %T\n", data1)
	fmt.Printf("tipe data data2 : %T\n", data2)

	// Number (float)

	/*
		Float merupakan tipe data numerik desimal pada bahasaGo.
		Secara umum float dibagi menjadi 2 jenis yaitufloat32 dan float34
	*/

	var decimalNumber float32 = 3.14

	fmt.Printf("decimal Number : %f\n", decimalNumber)
	fmt.Printf("decimal Number : %.3f\n", decimalNumber)

	/*
		Jika kita lihat pada gambar kedua, format angka desimal yang dicetak
		di terminal menghasil format angka yang berbeda walaupun
		masih berasal dari satu variable yang sama. Pada dasarnya verb %f
		digunakan untuk memformat angka desimal atau tipe data float
		menjadi 6 digit angka desimal.
	*/

	// Bool

	/*
		Tipe data bool terdiri dari 2 nilai yaitu false, dan true.
	*/

	var condition bool = true

	fmt.Printf("is it permitted? %t \n", condition)

	// String
	var message string = "halo"
	fmt.Println(message)

	var pesan = `Selamat datang di "Hacktiv8".
	Salam kenal :).
	Mari belajar "Scalable Web Service With Go"`
	fmt.Println(pesan)

	/*
		Semua tipe data yang sudah dibahas sebelumnya
		memiliki zero value (nilai default tipe data).
		Artinya meskipun variabeldideklarasikan dengan
		tanpa nilai awal, tetap akan ada nilai default-nya.
		- Zero value dari string adalah "" (string kosong).
		- Zero value dari bool adalah false.
		- Zero value dari tipe numerik non-desimal adalah 0.
		- Zero value dari tipe numerik desimal adalah 0.0.
	*/

	/*
		Zero value berbeda dengan nil.
		Nil adalah nilai kosong, benar-benar kosong.
		Nil tidak bisa digunakan pada tipe datayang
		sudah dibahas sebelumnya.
		Ada beberapa tipe data yang bisa di-set
		nilainya dengan nil, diantaranya:
		- pointer
		- tipe data function
		- slice
		- map
		- channel
		- interface kosong atau interface{}
	*/
}
