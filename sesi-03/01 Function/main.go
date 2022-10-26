package main

import (
	"fmt"
	"math"
	"strings"
)

/*
	Bahasa Go juga memiliki function atau fungsi nya sendiri. Cara menulis sebuah function pada Go cukup mudah. Caranya
adalah dengan menggunakan keyword func lalu diikuti dengan nama function dan parameter yang digunakan. Contohnya
seperti pada gambar pertama dibawah. Kita telah membuat function bernama greet yang menerima 2 parameter dengan tipe
data string dan int8. Function greet merupakan sebuah function yang tidak mengembalikan/me-return nilai apapun, melainkan
function ini hanya bertugas untuk menampilkan text pada terminal kita menggunakan fmt.Printf. Ketika kita jalankan pada
terminal kita, maka hasilnya akan seperti pada gambar kedua dibawah.
Perlu diingat disini bahwa, kita harus menjelaskan jenis dari tipe data nya pada parameter function jika function yang kita
gunakan menerima sebuah parameter.

*/

func main() {
	greet("Rizky", 25)

	greet2("Rizky", "Gorontalo")

	var names = []string{"Rizky", "Jordan"}
	var printMsg = greet3("Heii", names)

	fmt.Println(printMsg)

	var diameter float64 = 15
	var area, circumference = calculate(diameter)

	fmt.Println("Area: ", area)
	fmt.Println("Circumference: ", circumference)

	studentLists := print("Leno", "Eno", "Ovo", "Reza", "Yatman")

	fmt.Printf("%v\n", studentLists)

	numberLists := []int{1, 2, 3, 4, 5, 6, 7, 8}
	result := sum(numberLists...)
	fmt.Println("Result:", result)

	profile("Rizky", "pasta", "ayam geprek", "ikan roa", "sate padang")

}

//Functions

func greet(name string, age int8) {
	fmt.Printf("Hello there! My name is %s and I'm %d years old\n", name, age)
}

func greet2(name, address string) {
	fmt.Println("Hello there! My name is", name)
	fmt.Println("I live in", address)
}

// Function (Return)
/*
	Kita juga memberikan separator atau pemisah berupa spasi “ “ yang
	kita berikan pada argumen kedua pada function Join.

	Function Sprintf memiliki fungsi yang sama seperti function Printf,
	namun bedanya adalah function Sprintf akan me-return sebuah nilai
	sedang function Printf tidak
*/
func greet3(msg string, names []string) string {
	var joinStr = strings.Join(names, " ")
	var result string = fmt.Sprintf("%s %s", msg, joinStr)

	return result
}

// Function (Returning multiple values)
// Function pada bahasa Go dapat me-return lebih dari satu nilai.
/*
	Jika kita menggunakan teknik predefined return value, maka kita
perlu me-reassign variable yang digunakan sebagai nilai return
dengan sebuah nilai yang nantinya akan menghasilkan sebuah nilai
return. Tapi perlu diingat bahwa kita masih tetap membutuhkan
keyword return yang diletakkan di akhir baris function.
*/
func calculate(d float64) (float64, float64) {
	//Menghitung Luas
	var area float64 = math.Pi + math.Pow(d/2, 2)

	//Menghitung Keliling
	var circumference = math.Pi + d

	return area, circumference
}

// Function (Variadic function #1)
/*
	Bahasa Go telah menyediakan sebuah function yang dimana function
	tersebut dapat menerima argumen yang tak terbatas jumlahnya.
	Function tersebut bernama function variadic.
*/

/*
Function print merupakan sebuah function variadic. Cara
mengetahuinya adalah dengan memperhatikan penulisan pada
parameter yang diterimanya. Terdapat tanda titik sebanyak tiga kali
sebelum keterangan tipe data parameter yang diterimanya ...string.
*/
func print(names ...string) []map[string]string {
	var result []map[string]string

	for i, v := range names {
		key := fmt.Sprintf("Student%d", i+1)
		temp := map[string]string{
			key: v,
		}
		result = append(result, temp)
	}

	return result
}

// Function (Variadic function #2)
/*
Kita juga dapat menggunakan tipe data slice sebagai parameter dari
sebuah function variadic

Jika kita ingin memakai tipe data slice sebagai argumen untuk
sebuah function variadic, maka kita perlu menggunakan tanda ellipsis
atau tanda titik tiga... .

*/
func sum(numbers ...int) int {
	total := 0

	for _, v := range numbers {
		total += v
	}

	return total
}

// Function (Variadic function #3)
/*
	Kita dapat menggabungkan parameter biasa dengan parameter
	variadic pada sebuah function variadic. Namun perlu diingat disini
	bahwa parameter variadic perlu diletakkan di posisi akhir parameter
*/
func profile(name string, favFoods ...string) {
	mergeFavFoods := strings.Join(favFoods, ", ")

	fmt.Println("Hello There! I'm", name)
	fmt.Println("I really love to eat", mergeFavFoods)
}
