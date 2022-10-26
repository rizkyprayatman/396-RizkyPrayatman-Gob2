package main

import (
	"fmt"
	"reflect"
)

/*
Reflect digunakan untuk melakukan inspeksi variabel, mengambil informasi dari variabel tersebut atau bahkan
memanipulasinya.
Cakupan informasi yang bisa didapatkan lewat reflection sangat luas, seperti melihat struktur variabel, tipe, nilai pointer,
dan banyak lagi.

Go menyediakan package reflect, berisikan banyak sekali fungsi untuk keperluan reflection.
Dari banyak fungsi yang tersedia di dalam package tersebut, ada 2 fungsi yang paling penting untuk diketahui, yaitu
reflect.ValueOf() dan reflect.TypeOf().
● Fungsi reflect.ValueOf() akan mengembalikan objek dalam tipe reflect.Value, yang berisikan informasi yang
berhubungan dengan nilai pada variabel yang dicari
● Sedangkan reflect.TypeOf() mengembalikan objek dalam tipe reflect.Type. Objek tersebut berisikan informasi yang
berhubungan dengan tipe data variabel yang dicari
*/

// Idetifying Method Information
type student struct {
	Name  string
	Grade int
}

func (s *student) SetName(name string) {
	s.Name = name
}

func main() {
	//Identifying Data type & Value
	/*
		Dengan reflection, tipe data dan nilai variabel dapat
		diketahui dengan mudah.

		Fungsi reflect.valueOf() memiliki parameter yang bisa
		menampung segala jenis tipe data.

		Fungsi tersebut mengembalikan objek dalam tipe
		reflect.Value, yang berisikan informasi mengenai
		variabel yang bersangkutan.

		Objek reflect.Value memiliki beberapa method, salah
		satunya Type().

		Method ini mengembalikan tipe data variabel yang
		bersangkutan dalam bentuk string\

		Statement reflectValue.Int() menghasilkan nilai int dari
		variabel number.

		Untuk menampilkan nilai variabel reflect, harus
		dipastikan dulu tipe datanya.

		Ketika tipe data adalah int, maka bisa menggunakan
		method Int().

		Ada banyak lagi method milik struct reflect.Value yang
		bisa digunakan untuk pengambilan nilai dalam bentuk
		tertentu

		https://godoc.org/reflect#Kind
	*/
	var number = 23
	var reflectValue = reflect.ValueOf(number)

	fmt.Println("Tipe variabel: ", reflectValue.Type())

	if reflectValue.Kind() == reflect.Int {
		fmt.Println("Nilai variabel : ", reflectValue.Int())
	}

	//Accessing Value Using Interface Method
	/*
		Jika nilai hanya diperlukan untuk ditampilkan ke output, bisa menggunakan .Interface().
		Melalui method tersebut segala jenis nilai bisa diakses dengan mudah.
	*/
	fmt.Println("Tipe variabel: ", reflectValue.Type())
	fmt.Println("Nilai variabel: ", reflectValue.Interface())

	var nilai = reflectValue.Interface().(int)
	fmt.Println("Nilai: ", nilai)

	//Idetifying Method Information
	/*
		Informasi mengenai method juga bisa diakses lewat reflect, syaratnya masih sama seperti pada
		pengaksesan property, yaitu harus bermodifier public.
	*/
	var s1 = &student{Name: "John Wick", Grade: 2}
	fmt.Println("Nama: ", s1.Name)

	var reflectValue2 = reflect.ValueOf(s1)
	var method = reflectValue2.MethodByName("SetName")
	method.Call([]reflect.Value{
		reflect.ValueOf("Wick"),
	})

	fmt.Println("Nama: ", s1.Name)
}
