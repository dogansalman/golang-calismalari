package main

import (
	"fmt"
)

func main(){

	var i int32
	fmt.Println("Sayı gir")
	fmt.Scan(&i)

	switch i{
		case 0: fmt.Println("Sıfır")
		case 1: fmt.Println("Bir")
		case 2: fmt.Println("İki")
		default: fmt.Println("Bilinmeyen Sayı")
	}

}
