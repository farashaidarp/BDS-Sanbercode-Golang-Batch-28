package main

import (
	"fmt"
)

func main() {

	//soal 1
	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			fmt.Println("berkualitas")
		} else if i%3 == 0 {
			fmt.Println("I LOVE CODING")
		} else {
			fmt.Println("Santai")
		}

	}
	//soal2
	for i := 1; i <= 7; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	//soal3
	var kalimat = []string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}
	var kalimats = kalimat[2:7]
	fmt.Println(kalimats)

	//soal 4
	var sayuran = []string{"Bayam", "Buncis", "Kangkung", "Kubis", "seledri", "Tuage", "Timun"}
	for i := 0; i < 7; i++ {
		fmt.Println(i+1, sayuran[i])
	}

	//soal 5
	var satuan map[string]int
	satuan = map[string]int{}

	satuan["panjang"] = 7
	satuan["lebar"] = 4
	satuan["tinggi"] = 6

	fmt.Println("Panjang = ", satuan["panjang"])
	fmt.Println("lebar = ", satuan["lebar"])
	fmt.Println("tinggi = ", satuan["tinggi"])
	fmt.Println("Volume = ", satuan["panjang"]*satuan["lebar"]*satuan["tinggi"])

}
