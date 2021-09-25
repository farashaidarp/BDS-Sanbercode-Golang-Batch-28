package main

import 
	"fmt"

func luasPersegiPanjang(angka1 ,angka2 int) int{
	return angka1 * angka2
	}
func kelilingPersegiPanjang(angka1 , angka2  int)int{
	return 2 *(angka1+angka2)
	}
func volumeBalok(angka1,angka2,angka3 int)int{
	return angka1*angka2*angka3
	}

func introduce(nama1, nama2, nama3 ,nama4 string) (hasil string) {
	hasil = "Pak "+ nama1 + " adalah seorang " + nama3 +" " +  nama2 +" yang berusia "+ nama4 +" tahun"
	return
	}
func introduce1(nama1, nama2, nama3 ,nama4 string) (hasil string) {
	hasil = "Bu "+ nama1 + " adalah seorang " + nama3 +" " +  nama2 +" yang berusia "+ nama4 +" tahun"
	return
	}	
	func buahFavorit(name string, buah ...string) {
		fmt.Printf("halo " + " nama saya " + name + " dan buah favorit saya adalah ")
		for _, buah := range buah {
			fmt.Printf(buah + " ")
		}
	
	}
	
func main() {

	//soal 1
	panjang := 12
	lebar := 4
	tinggi := 8

	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, lebar)
	volume := volumeBalok(panjang, lebar, tinggi)

	fmt.Println(luas) 	
	fmt.Println(keliling)
	fmt.Println(volume)

	//soal2
	john := introduce("John", "laki-laki", "penulis", "30")
	fmt.Println(john) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"

	sarah := introduce1("Sarah", "perempuan", "model", "28")
	fmt.Println(sarah) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"

	//soal3


	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}

	 buahFavorit("John",buah...)

	
	// halo nama saya john dan buah favorit saya adalah "semangka", "jeruk", "melon", "pepaya"*/

	//soal 4
	



}