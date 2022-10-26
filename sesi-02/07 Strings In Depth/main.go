package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	/*
		Tipe data string pada Go terbentuk dari kumpulan tipe data byte yang di letakkan di dalam slice atau bisa kita sebut
		dengan slice of bytes.
		Tipe data byte pada Go merupakan tipe data alias dari tipe data uint8.
		Ketika kita melakukan indexing terhadap suatu string, maka kita akan mendapat nilai representasi dari byte nya
	*/

	// Looping Over String
	// Ketika kita melakukan indexing terhadap suatu string, maka kita akan mendapat nilai representasi dari byte nya.
	city := "Jakarta"
	for i := 0; i < len(city); i++ {
		fmt.Printf("index: %d, byte: %d\n", i, city[i])
	}

	/*
		Pada looping tersebut, kita melakukannya dengan pendekatan byte-per-byte, atau dapat diartikan kita melooping
		string “Jakarta” berdasarkan byte nya. Kode city[i] tidak akan menghasilkan karakternya, melainkan byte nya.
	*/

	var city2 []byte = []byte{74, 97, 107, 97, 114, 116, 97}
	fmt.Println(string(city2))

	str1 := "makan"

	str2 := "mânca"

	fmt.Printf("str1 byte length => %d\n", len(str1))
	fmt.Printf("str2 byte length => %d\n", len(str2))

	fmt.Printf("str1 character length => %d\n", utf8.RuneCountInString(str1))
	fmt.Printf("str2 character length => %d\n", utf8.RuneCountInString(str2))

	fmt.Println("\n\nLooping String")
	for i, s := range str2 {
		fmt.Printf("index => %d, string => %s\n", i, string(s))
	}

}
