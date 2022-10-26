package main

import "fmt"

func main() {
	// Conditions(Temporary variable)

	// If Else

	// var currentYear = 2022

	// var age = currentYear - 1998

	// if age < 17 {
	// 	fmt.Println(("Kamu belum boleh membuat kartu sim"))
	// } else {
	// 	fmt.Println(("Kamu sudah boleh membuat kartu sim"))
	// }

	var currentYear = 2022

	if age := currentYear - 1998; age < 17 {
		fmt.Println(("Kamu belum boleh membuat kartu sim"))
	} else {
		fmt.Println(("Kamu sudah boleh membuat kartu sim"))
	}

	// Switch

	var score = 8

	switch score {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	// Switch with relational operators

	var nilai = 7

	switch {
	case nilai == 8:
		fmt.Println("Sempurna")
	case (nilai < 8) && (nilai > 3):
		fmt.Println("Cukup")
	default:
		{
			fmt.Println("Yok Belajar, Semangat")
			fmt.Println("Kamu butuh belajar lebih banyak")
		}
	}

	// Switch (fallthrough keyword)

	var ipk = 3.0

	switch {
	case ipk == 4.0:
		fmt.Println("Sempurna")
	case (ipk < 3.5) && (ipk > 2.5):
		fmt.Println("memuaskan")
		fallthrough
	case ipk < 2:
		fmt.Println("Sudah baik, tapi tetap semangat belajar ya")
	default:
		{
			fmt.Println("Yok Belajar, Semangat")
			fmt.Println("Kamu butuh belajar lebih banyak")
		}
	}

	// Nested Conditions

	var jumlah = 8

	if jumlah > 7 {
		switch jumlah {
		case 10:
			fmt.Println("Komplit")
		default:
			fmt.Println("Kurang")
		}
	} else {
		if jumlah == 5 {
			fmt.Println("Tambah lagi 5")
		} else if jumlah == 3 {
			fmt.Println("Tambah lagi 7")
		} else {
			fmt.Println("Semangat Tambah lagi")
			if jumlah == 0 {
				fmt.Println("Kurang banget, tambah lagi")
			}
		}
	}

}
