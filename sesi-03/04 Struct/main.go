package main

import (
	"fmt"
	"strings"
)

/*
	Bahasa Go tidak menyediakan tipe data class untuk membuat
	sebuah struktur data seperti layaknya bahasa pemrograman yang
	menganut paradigma OOP. Tetapi Go menyediakan sebuah tipe
	data yang bernama Struct untuk membuat sebuah struktur.

	Struct adalah sebuah tipe data berupa kumpulan/koleksi dari
	berbagai macam property/field dan juga method. Cara
	membuat struct pada Go cukup mudah.

	Untuk membuat sebuah tipe data struct , kita perlu membuat
	terlebih dahulu strukturnya dengan urutan format penulisan:
	● Penulisan keyword type
	● Nama dari struct
	● Penulisan keyword struct
	● Kemudian diikuti dengan tanda kurung kurawal {}
	● Mendefinisikan field yang diinginkan
*/

// Struct
type Employee struct {
	name     string
	age      int
	division string
}

type Person struct {
	name string
	age  int
}

type Employee2 struct {
	division string
	person   Person
}

func main() {
	//Giving value to struct
	/*
		Agar kita dapat memberikan nilai kepada field yang terdapat
		pada sebuah struct, kita perlu terlebih dahulu menyimpan tipe
		data dari struct kepada sebuah variable
	*/
	var employee Employee

	employee.name = "Rizky"

	employee.age = 25

	employee.division = "Fullstack Web"

	fmt.Println(employee.name)
	fmt.Println(employee.age)
	fmt.Println(employee.division)

	//Initializing Struct
	/*
		Kita juga dapat menginisialisasi sebuah struct sekaligus
		memberikan nilai-nilai nya.
	*/
	var employee1 = Employee{}
	employee1.name = "Reza"
	employee1.age = 22
	employee1.division = "UI Design"

	var employee2 = Employee{name: "Kity", age: 22, division: "IT Support"}

	fmt.Printf("Employee1: %+v\n", employee1)
	fmt.Printf("Employee2: %+v\n", employee2)

	//Pointer to a struct
	// Kita juga dapat menggunakan pointer pada sebuah struct

	var employee3 = Employee{name: "Eno", age: 22, division: "Tech Nano"}

	var employee4 *Employee = &employee3

	fmt.Println("Employee3 name:", employee3.name)
	fmt.Println("Employee4 name:", employee4.name)

	fmt.Println(strings.Repeat("#", 20))

	employee4.name = "Reza"

	fmt.Println("Employee3 name:", employee3.name)
	fmt.Println("Employee4 name", employee4.name)

	//Embedded struct
	/*
		Struct juga dapat mengandung tipe data struct lainnya dengan
		menjadikannya sebuah field.
	*/
	var employee5 = Employee2{}

	employee5.person.name = "Lenov"
	employee5.person.age = 17
	employee5.division = "Electrician"

	fmt.Printf("%+v \n", employee5)

	//Anonymous Struct
	/*
		Anonymous Struct tanpa pengisian field

		Anonymous struct adalah sebuah struct yang tidak
		dideklerasikan di awal sebagai sebuah tipe data struct baru,
		melainkan langsung dideklerasikan bersamaan dengan
		pembuatan variable.
	*/

	var employee6 = struct {
		person   Person
		division string
	}{}
	employee6.person = Person{name: "Nov", age: 28}
	employee6.division = "Mobile"

	//Anonymous Struct dengan pengisian field
	var employee7 = struct {
		person   Person
		division string
	}{
		person:   Person{name: "Deni", age: 20},
		division: "Web",
	}

	fmt.Printf("Employee6: %+v\n", employee6)
	fmt.Printf("Employee7: %+v\n", employee7)

	//Slice of Struct
	/*
		Slice juga dapat dikombinasikan dengan tipe data struct, cara
		penulisannya mirip seperti slice of map yang telah kita bahas
		pada materi sebelumnya.
	*/
	var people = []Person{
		{name: "Airell", age: 23},
		{name: "Amanda", age: 23},
		{name: "Mailo", age: 23},
	}

	for _, v := range people {
		fmt.Printf("%+v\n", v)
	}

	//Slice of anonymous struct
	/*
		Anonymous struct juga dapat digabungkan dengan tipe data
		slice dan pengisian nilainya pun dapat dilakukan secara
		langsung.
	*/
	var employee8 = []struct {
		person   Person
		division string
	}{
		{person: Person{name: "Airell", age: 23}, division: "Curriculum Developer"},
		{person: Person{name: "Ananda", age: 23}, division: "Finance"},
		{person: Person{name: "Mailo", age: 21}, division: "Marketing"},
	}

	for _, v := range employee8 {
		fmt.Printf("%+v\n", v)
	}
}
