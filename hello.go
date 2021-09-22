package main

import (
	"fmt"
	"strings"
)

func main() {

	// contoh 1
	var text2 string = "ini teks 2"
	text2 = "ini teks 1 diubah"
	fmt.Println(text2)

	var decimalNumber = 2.62

	fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("bilangan desimal: %.2f\n", decimalNumber)

	var exist bool = true
	fmt.Printf("exist? %t \n", exist)

	var message string = "Halo"
	fmt.Printf("message: %s \n", message)

	var index1 = strings.Index("ethan hunt", "nt")
	fmt.Println(index1)

	var text = "bananana"
	var find = "a"
	var replaceWith = "o"

	var newText1 = strings.Replace(text, find, replaceWith, 1)
	fmt.Println(newText1) // "bonana"

	var newText2 = strings.Replace(text, find, replaceWith, 2)
	fmt.Println(newText2) // "bonona"

	var newText3 = strings.Replace(text, find, replaceWith, -1)
	fmt.Println(newText3) // "bonono"
}
