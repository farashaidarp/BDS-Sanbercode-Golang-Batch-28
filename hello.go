package main

import "fmt"

func main() {
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

	


}
