package main

import "fmt"

func main() {

	/*
		Ketika kita ingin membuat sebuah map, maka kita juga harus
		langsung menginisialisasikannya.

		Penulisan map[string]string berarti tipe data key dari map harus
		string dan value dari map juga harus string. Lalu jika kita ingin
		memberikan nilai atau value kepada map nya, maka harus di
		inisialisasi terlebih dahulu.

	*/

	var person map[string]string // Deklarasi
	person = map[string]string{} // Inisialisasi

	person["name"] = "Rizky"
	person["age"] = "25"
	person["address"] = "Jalan Rambutan"

	fmt.Println("Name: ", person["name"])
	fmt.Println("Age: ", person["age"])
	fmt.Println("Address: ", person["address"])

	// Map With Value
	/*
		Kita juga dapat memberikan value beserta dengan key nya dengan
		langsung.

	*/
	var person2 = map[string]string{
		"name":    "Rizky",
		"age":     "25",
		"address": "Jalan Rambutan",
	}
	fmt.Println("Name: ", person2["name"])
	fmt.Println("Age: ", person2["age"])
	fmt.Println("Address: ", person2["address"])

	// Looping Map
	/*
		Ketika kita ingin melakukan loop terhadap sebuah map, maka kita
		dapat menggunakan range loop.

	*/
	var person3 = map[string]string{
		"name":    "Rizky",
		"age":     "25",
		"address": "Jalan Rambutan",
	}

	for key, value := range person3 {
		fmt.Println(key, ":", value)
	}

	// Delete Map
	/*
		Kita juga dapat menghapus value dari sebuah map dengan cara
		menggunakan fungsi delete

		Argumen pertama yang kita berikan pada fungsi
		delete adalah map yang ingin kita hapus atau bisa juga variable
		tempat map disimpan. Lalu argumen kedua adalah key yang
		menyimpan value nya.
	*/
	fmt.Println("Before Deleting", person3)
	delete(person3, "age")
	fmt.Println("After Deleting", person3)

	// Map ( Detectng a value )
	/*
		Kita juga dapat mendeteksi keberadaan suatu value dengan cara
		mengakses key dari map nya lalu memberikan 2 variable sebagai
		penampungnya.

		Variable value yang kita berikan akan mengembalikan value asli dari
		map nya jika memang key nya ada, jika tidak ada maka kita akan
		mendapat zero value dari value nya. Kemudian variable exist yang
		kita berikan akan mengembalikan nilai dengan tipe data bool.
		Variable exist akan mengembalikan true jika memang key nya ada,
		jika tidak maka akan mengembalikan false.
	*/
	var person4 = map[string]string{
		"name":    "Rizky",
		"age":     "25",
		"address": "Jalan Rambutan",
	}

	value, exist := person4["age"]

	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("value does'nt exist")
	}

	var person5 = map[string]string{
		"name":    "Rizky",
		"age":     "25",
		"address": "Jalan Rambutan",
	}

	delete(person5, "age")

	value, exist = person5["age"]

	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("value has been deleted")
	}

	// Map ( Combining slice )
	/*
		Kita juga dapat mengkombinasikan slice dengan map.
		Contohnya seperti pada gambar pertama di sebelah kanan.
		Kita juga dapat me-looping nya dengan menggunakan range
		loop.
	*/
	fmt.Println("\n\nMap ( Combining slice )")
	var people6 = []map[string]string{
		{"name": "Lenov", "age": "20"},
		{"name": "Eno", "age": "28"},
		{"name": "Ovo", "age": "22"},
	}
	for i, person6 := range people6 {
		fmt.Printf("Index : %d, name: %s, age: %s\n", i, person6["name"], person6["age"])
	}

}
