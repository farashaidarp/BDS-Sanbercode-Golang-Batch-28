package main

import (
	"fmt"
)
//soal 1

func keliling(data float64 , jari float64) *float64 {
	data = (jari  +  jari) * 22/7
	return &data
}
func luas(data float64 , jari float64) *float64 {
	data = (jari * jari) * 22/7
	return &data
}

func introduce(hasil ,nama1, nama2, nama3 ,nama4 string) *string {
	hasil = "Pak "+ nama1 + " adalah seorang " + nama3 +" " +  nama2 +" yang berusia "+ nama4 +" tahun"
	return &hasil
	}
func introduce1(hasil,nama1, nama2, nama3 ,nama4 string) * string{
	hasil = "Bu "+ nama1 + " adalah seorang " + nama3 +" " +  nama2 +" yang berusia "+ nama4 +" tahun"
	return &hasil
	}	

func main() {

	var a float64 
	var c float64 = 10

	
	luaslingkaran := luas(a ,c)
	kelilinglingkaran := keliling(a,c)


	fmt.Println("Luas lingkaran adalah", *luaslingkaran)
	fmt.Println("Luas lingkaran adalah", *kelilinglingkaran)

	//soal 2

	var b string 

	john := introduce(b ,"John", "laki-laki", "penulis", "30")
	fmt.Println(*john) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"

	sarah := introduce1(b, "Sarah", "perempuan", "model", "28")
	fmt.Println(*sarah) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"

	//soal 3
	var sayuran = []string{"Bayam", "Buncis", "Kangkung", "Kubis", "seledri", "Tuage", "Timun"}
	p := &sayuran

	fmt.Println("Element 0: ", (*p)[0])

	fmt.Println("List of Elements")
	for i := 0; i < len(sayuran); i++ {
		fmt.Print((*p)[i], "  ")
	}

	fmt.Println("\nList of Elements")
	for index, value := range *p {
		fmt.Println(index + 1, value)
	}

	//soal 4

	
	}
