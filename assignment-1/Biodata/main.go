package main

import (
	"fmt"
	"os"
	"strconv"
)

type Student struct {
	noAbsen   string
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

func main() {
	var argsRaw = os.Args
	var args = argsRaw[1]
	noAbsen, err := strconv.Atoi(args)
	_ = err

	getNoAbsen(noAbsen)
}

func getNoAbsen(noAbsen int) {
	var studentList = []Student{
		{noAbsen: "1", nama: "Dewa", alamat: "Jalan Kh. Adam Zakaria", pekerjaan: "Data Analitik", alasan: "Menyukai data"},
		{noAbsen: "2", nama: "Julai", alamat: "Jalan Kh. Adam Zakaria", pekerjaan: "Operator Gudang", alasan: "suka ngarsip dan merapikan apapun"},
		{noAbsen: "3", nama: "Haris", alamat: "Jalan Sirsak", pekerjaan: "Owner Rumah Bersih", alasan: "Bersih itu sehat dan menyehatkan"},
		{noAbsen: "4", nama: "Murai", alamat: "Perum Awarta", pekerjaan: "Pembuat Merchandise", alasan: "Merchandise dapat membuat orang senang"},
		{noAbsen: "5", nama: "Advendy", alamat: "Jalan Jendral Sudirman", pekerjaan: "Service Elektronik", alasan: "Alat elektronik butuh perhatian"},
		{noAbsen: "6", nama: "Egata", alamat: "Jalan Rambutan", pekerjaan: "Pelukis", alasan: "Seni merupakan wujud apresiasi segala penciptaan Tuhan"},
		{noAbsen: "7", nama: "Beltik", alamat: "Jalan Palma", pekerjaan: "Pengrajin", alasan: "Asyik dapat menyalurkan isi pikiran"},
		{noAbsen: "8", nama: "Reyden", alamat: "Jalan Mangga", pekerjaan: "Front End Web", alasan: "Segala bentuk keindahan sebuah web merupakan motivasi"},
		{noAbsen: "9", nama: "Keyray", alamat: "Jalan Jakarta", pekerjaan: "Art Digital Desain", alasan: "Digitalisasi mengurangi dampak lingkungan"},
		{noAbsen: "10", nama: "Rezka", alamat: "Jalan Kalimantan", pekerjaan: "Back End Web", alasan: "Dapat menjadi Penyambung tangan "},
	}

	switch {
	case noAbsen <= 10:
		fmt.Println("Student =>")
		fmt.Println("No Absen  :", studentList[noAbsen-1].noAbsen)
		fmt.Println("Nama      :", studentList[noAbsen-1].nama)
		fmt.Println("Alamat    :", studentList[noAbsen-1].alamat)
		fmt.Println("Pekerjaan :", studentList[noAbsen-1].pekerjaan)
		fmt.Println("Alasan    :", studentList[noAbsen-1].alasan)
	default:
		fmt.Println("Data yang dicari tidak ditemukan!!")
	}
}
