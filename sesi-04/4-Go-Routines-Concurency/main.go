package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	/*
		Arti dari concurrency adalah mengeksekusi sebuah proses secara independen atau berhadapan dengan lebih dari satu tugas
		dalam waktu yang sama. Perlu diingat disini bahwa concurrency berbeda dengan parallelism, karena parallelism memiliki arti
		mengerjakan tugas yang banyak secara bersamaan dari awal hingga akhir. Sedangkan pada concurrency, kita tidak akan tahu
		tentang siapa yang akan menyelesaikan tugasnya terlebih dahulu.
	*/
	fmt.Println("main execution started")

	/*
		Goroutine merupakan sebuah thread ringan pada bahasa Go untuk
		melakukan concurrency. Jika dibandingkan dengan thread biasa,
		Goroutine memiliki ukuran thread yang jauh lebih ringan. Pada saat
		kita mengeksekusi sebuah Goroutine, maka satu Goroutine hanya
		membutuhkan 2kb memori saja, sedangkan satu thread biasa dapat
		menghabiskan 1-2mb memori.
		Goroutine bersifat asynchronous sehingga proses nya tidak saling
		tunggu dengan Goroutine lainnya.
		Untuk membuat sebuah Goroutine, maka kita harus terlebih dahulu
		membuat sebuah function. Lalu ketika kita ingin memanggil function
		tersebut, maka kita perlu menggunakan keyword go sebelum kita
		memanggil function tersebut
	*/

	go firstProcess(8)
	go secondProcess(8)

	fmt.Println("no. of go routines : ", runtime.NumGoroutine())

	time.Sleep(time.Second * 2) //eksekuti timeout

	fmt.Println("main execution ended")
}

func firstProcess(index int) {
	fmt.Println("first process func started")
	for i := 0; i < index; i++ {
		fmt.Println("i=", i)
	}
	fmt.Println("first process func ended")
}
func secondProcess(index int) {
	fmt.Println("second process func started")
	for j := 0; j < index; j++ {
		fmt.Println("j=", j)
	}
	fmt.Println("Second process func ended")
}
