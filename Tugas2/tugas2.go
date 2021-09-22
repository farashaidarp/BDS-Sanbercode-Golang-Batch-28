package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// Soal 1
	var message = `Bootcamp `
	var message1 = `Digital `
	var message2 = `Skill`
	var message3 = `Sanbercode `
	var message4 = `Golang `
	fmt.Println(message, message1, message2, message3, message4)

	//soal 2
	halo := "Halo Dunia"
	var find = "Dunia"
	var ubah = "Golang"
	var ganti = strings.Replace(halo, find, ubah, 1)
	fmt.Println(ganti)

	//soal3
	var kataPertama = "saya"
	var kataKedua = "senang"
	var kataKetiga = "belajar"
	var kataKeempat = "golang"

	var ubah1 = strings.Replace(kataKedua, "s", "S", 1)
	var ubah2 = strings.Replace(kataKetiga, "r", "R", 1)
	var ubah3 = strings.ToUpper(kataKeempat)

	fmt.Println(kataPertama, ubah1, ubah2, ubah3)

	//soal4
	var angkaPertama = "8"
	var angkaKedua = "5"
	var angkaKetiga = "6"
	var angkaKeempat = "7"

	var angka1, err = strconv.Atoi(angkaPertama)
	var angka2, ere = strconv.Atoi(angkaKedua)
	var angka3, ert = strconv.Atoi(angkaKetiga)
	var angka4, erw = strconv.Atoi(angkaKeempat)

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
	fmt.Println(angka1 + angka2 + angka3 + angka4)

	//soal 5
	kalimat := "halo halo bandung"
	angka := 2021

	var rep = strings.ReplaceAll(kalimat, "halo", "Hi")

	fmt.Println(rep, "-", angka)
}
