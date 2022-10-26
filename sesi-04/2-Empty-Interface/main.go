package main

import (
	"fmt"
)

func main() {
	//Empty Interface
	/*
		Empty interface merupakan suatu tipe data yang dapat
		menerima segala tipe data pada bahasa Go. Maka dari itu,
		sebuah variable dengan tipe data empty interface dapat
		diberikan nilai dengan tipe data yang berbeda-beda.
	*/
	var randomValues interface{}

	_ = randomValues

	randomValues = "Jalan Sudirman"

	randomValues = 20

	randomValues = true

	randomValues = []string{"Airell", "Nanda"}

	fmt.Println(randomValues)

	//Empty interface (Type assertion)
	var v interface{}

	v = 20

	if value, ok := v.(int); ok == true {
		v = value * 9
	}
	fmt.Println(v)

	//Empty interface (Map & Slice with Empty Interface)
	rs := []interface{}{1, "Airiell", true, 2, "Nanda", true}

	rm := map[string]interface{}{
		"Name":   "Airiell",
		"Status": true,
		"Age":    23,
	}

	_, _ = rs, rm
}
