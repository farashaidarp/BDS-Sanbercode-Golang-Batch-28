package main

import (
	"fmt"
	"strconv"
)

func main() {

	//soal 1

	var panjangPersegiPanjang string = "8"
	var lebarPersegiPanjang string = "5"
	var alasSegitiga string = "6"
	var tinggiSegitiga string = "7"

	var angka1, err = strconv.Atoi(panjangPersegiPanjang)
	var angka2, ere = strconv.Atoi(lebarPersegiPanjang)
	var angka3, ert = strconv.Atoi(alasSegitiga)
	var angka4, erw = strconv.Atoi(tinggiSegitiga)

	if err == nil {
		fmt.Println(angka1)
	}
	if ere == nil {
		fmt.Println(angka2)
	}
	if ert == nil {
		fmt.Println(angka3)
	}
	if erw == nil {
		fmt.Println(angka4)
	}
	var luasPersegiPanjang int = angka1 * angka2
	var kelilingPersegiPanjang int = 2 * (angka1 + angka2)
	var luasSegitiga int = (angka3 * angka4) / 2
	fmt.Println(luasPersegiPanjang)
	fmt.Println(kelilingPersegiPanjang)
	fmt.Println(luasSegitiga)

	//soal 2

	var nilaiJohn = 80
	var nilaiDoe = 50

	if nilaiJohn >= 80 {
		fmt.Println("joe dapat A")
	} else if nilaiJohn >= 70 && nilaiJohn < 80 {
		fmt.Println("joe Dapat b")
	} else if nilaiJohn >= 60 && nilaiJohn < 70 {
		fmt.Println("Joe dapat nilai C")
	} else if nilaiJohn >= 50 && nilaiJohn < 60 {
		fmt.Println("Joe dapat nilai D")
	} else {
		fmt.Println("Joe dapat nilai E")
	}

	if nilaiDoe >= 80 {
		fmt.Println("doe Dapat A")
	} else if nilaiDoe >= 70 && nilaiDoe < 80 {
		fmt.Println("Doe Dapat b")
	} else if nilaiDoe >= 60 && nilaiDoe < 70 {
		fmt.Println("Doe dapat nilai C")
	} else if nilaiDoe >= 50 && nilaiDoe < 60 {
		fmt.Println("Doe dapat nilai D")
	} else {
		fmt.Println("Doe dapat nilai E")
	}

	//soal 3

	var tanggal = 22
	var bulan = 7
	var tahun = 2000

	switch bulan {
	case 1:
		fmt.Println(tanggal, "Januari", tahun)
	case 2:
		fmt.Println(tanggal, " February ", tahun)
	case 3:
		fmt.Println(tanggal, " Maret ", tahun)
	case 4:
		fmt.Println(tanggal, " April ", tahun)
	case 5:
		fmt.Println(tanggal, " Mei ", tahun)
	case 6:
		fmt.Println(tanggal, " Juni ", tahun)
	case 7:
		fmt.Println(tanggal, "Juli", tahun)
	case 8:
		fmt.Println(tanggal, " Agustus ", tahun)
	case 9:
		fmt.Println(tanggal, " Sepetember ", tahun)
	case 10:
		fmt.Println(tanggal, " Oktober ", tahun)
	case 11:
		fmt.Println(tanggal, " November ", tahun)
	case 12:
		fmt.Println(tanggal, " Desember ", tahun)
	default:
		fmt.Println("error")
	}

	//soal 4
	var lahir = 2000

	if lahir >= 1994 && lahir <= 1964 {
		fmt.Println("baby Boomer")
	} else if lahir >= 1965 && lahir <= 1979 {
		fmt.Println("Generasi X")
	} else if lahir >= 1980 && lahir < 1994 {
		fmt.Println("Generasi Y")
	} else if lahir >= 1995 && lahir <= 2015 {
		fmt.Println("Generasi Z")
	} else {
		fmt.Println("anda lahir diatas 2015")
	}

}
