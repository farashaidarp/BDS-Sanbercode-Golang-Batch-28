package main

import "fmt"

func buahFavorit(name string, buah ...string) {
 fmt.Printf("halo " + " nama saya " + name + " dan buah favorit saya adalah ")
	for _, buah := range buah {
		fmt.Printf(buah + " ")
	}
}

func main() {
	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}
	buahFavorit("John", buah...)
	// soal 4
}