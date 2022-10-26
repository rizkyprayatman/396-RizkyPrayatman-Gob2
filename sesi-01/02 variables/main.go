package main

import "fmt"

func main() {
	fmt.Println("Variables")
	fmt.Println("==============")

	/*
		Pada bahasa Golang sendiri, terdapat 2 cara penulisan untuk menulis variable.
		Yang pertama adalah yang dituliskan tipe datanya
		dan ada juga yang tidak ditulis tipe datanya atau biasa.
	*/

	// Variable with data type

	var name string = "Rizky"
	var age int = 24

	fmt.Println("nama saya", name)
	fmt.Println("umur saya", age)

	/*
		Kita bisa me-reassign suatu variable dengan suatu nilai
		asalkan kita memberikan nilai dengan tipe data yang sama
		dengan tipe data yang sudah dimiliki oleh variable tersebut.
	*/

	var sekolah string

	sekolah = "Universitas Negeri Gorontalo"

	fmt.Println("Sekolah di", sekolah)
	fmt.Println("-----------------------------")

	// Variable without data type

	/*
		Kita juga dapat membuat suatu variable dengan
		tidak memberikan tipe datanya secara eksplisit.
		Hal ini dikenal dengan nama type inference yang
		dimana sistem dari golang yang akan menentukan
		sendiri tipe datanya secara otomatis tergantung dari
		nilai yang kita berikan pada variable tersebut.
	*/

	var course = "Pemograman Golang"
	var price = 255

	fmt.Printf("%T, %T", course, price)
	fmt.Print("\n")

	// Variable without data type(Short Declaration)

	/*
		kita menggunakan tanda titik dua (:) sebelum
		tanda sama dengan (=).
		Dengan seperti ini maka kita telah menerapkan
		teknik short declaration yang dimana teknik
		ini cukup mempermudah kita dalam membuat suatu variable.
	*/

	pelajaran := "Fisika"
	kelas := 12

	fmt.Printf("%T, %T", pelajaran, kelas)
	fmt.Print("\n")

	fmt.Println("-----------------------------")

	// Multiple variable declarations

	var student1, student2, student3 string = "Rizky", "Nano", "Har"
	var nostudent1, nostudent2, nostudent3 int
	nostudent1, nostudent2, nostudent3 = 1, 2, 3

	fmt.Println(student1, student2, student3)
	fmt.Println(nostudent1, nostudent2, nostudent3)

	fmt.Println("-------------------------------------------------------")

	// Multiple variable declarations with Short Declaration

	var apotek, nokontak, alamat = "Apotek Sehati", 821, "Jalan Sudirman"
	obat, jumlah, jenis := "Paracetamol", 3, "Kapsul"

	fmt.Println(apotek, nokontak, alamat)
	fmt.Println(obat, jumlah, jenis)
	fmt.Println("-------------------------------------------------------")

	// Underscore variable
	/*
		Bahasa Golang memiliki satu aturan yang cukup strict/ketat,
		yang dimana tidak ada variable yang boleh menganggur ketika sudah kita buat.
	*/

	/*
		Dengan variable underscore maka kita dapat menghindari error yang akan
		terjadi jika kita mempunyai suatu variable yang menganggur atau tidak dipakai.
	*/

	var os string

	var laptop, unit, tipe = "lenovo", 1, "x220"

	_, _, _, _ = os, laptop, unit, tipe

	// fmt.Printf Usage

	/*
		fmt.Printf tidak akan membuat baris baru tidak seperti fmt.Println
		yang menambah baris baru secara otomatis.
	*/

	/*
		jika kita ingin mengetahui tipe data dari suatu variable,
		maka kita dapat menggunakan verb% T.
	*/

	fmt.Printf("Tipe data variable Laptop adalah %T \n", laptop)
	fmt.Printf("Laptop %s, jumlah unit %d, tipe %s", laptop, unit, tipe)
}
