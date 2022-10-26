package main

import (
	"fmt"
	"sync"
)

func main() {
	/*
		Waitgroup merupakan sebuah struct dari package sync, yang digunakan untuk melakukan sinkronisasi terhadap go
		routine
	*/
	fruits := []string{"apple", "mangga", "durian", "rambutan"}

	var wg sync.WaitGroup

	/*
		Method Add digunakan untuk menambah counter dari waitgroup.
		Counter dari waitgroup ini berguna untuk memberitahu waitgroup
		tentang jumlah go routine yang harus ditunggu.
	*/

	for index, fruit := range fruits {
		wg.Add(1)
		go printFruits(index, fruit, &wg)
	}
	wg.Wait()
	/*
		Method Wait berguna untuk menunggu seluruh go routine menyelesaikan proses nya,
		sehingga method Wait akan menahan function main hingga seluruh
		proses go routine selesai.
	*/
}

func printFruits(index int, fruit string, wg *sync.WaitGroup) {
	fmt.Printf("index => %d, fruit => %s\n", index, fruit)
	wg.Done()
	/*
		untuk memberitahu waitgroup tentang go routine yang telah
		menyelesaikan proses nya, maka kita perlu memanggil method Done
	*/
}
