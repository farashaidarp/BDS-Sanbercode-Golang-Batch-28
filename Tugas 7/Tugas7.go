package main

import (
	"fmt"
)

type  buah struct{
	nama string
	warna  string
	biji bool
	harga int
}
type segitiga struct{
	alas, tinggi int
  }
  
  type persegi struct{
	sisi int
  }
  
  type persegiPanjang struct{
	panjang, lebar int
  }

func (s segitiga) Hitungsegi3() int {
  return (s.alas * s.tinggi)/2
}

func (p persegi) Hitungsegi4() int {
	return p.sisi * p.sisi
  }

  func (pp persegiPanjang) Hitungsegipanjang() int {
	return pp.panjang * pp.lebar
  }


func main() {

	//soal 1
	fmt.Println("===== Soal 1 =====")
	fruit := [] buah {
		{
	nama:"Nanas",
	warna:"Kuning",
	biji : false,
	harga: 9000,
		},
		{
	nama:"jeruk",
	warna:"Orange",
	biji : true,
	harga: 8000,
				},
				{
	nama:"Semangka",
	warna:"Hijau & merah",
	biji : true,
	harga: 10000,
				},
			{
	nama:"Pisang",
	warna:"Kuning",
	biji : false,
	harga: 5000,
				},
	}

	fmt.Println( fruit)

	//soal 2

	   r := segitiga {10,21}
	   p := persegi {10}
	   pp := persegiPanjang{10,12}
	   fmt.Println(r.Hitungsegi3())
	   fmt.Println(p.Hitungsegi4())
	   fmt.Println(pp.Hitungsegipanjang())
	
}