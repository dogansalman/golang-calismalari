package  main

import "fmt"

func main(){

	// Array tanımlamaları

	x := [5]float64{ 98, 22, 31, 91, 7 }

	/*
		var x [5]float64
		x[0] = 65
		x[1] = 66
		x[2] = 67
		x[3] = 68
		x[4] = 69


		z := [5]float64{
		98,
		22,
		31,
		91,
		7,
		}
	*/


	/*
	Eğer array e bir boyut vermek istemiyorsak "make" kullanılıyor
	*/
	//d := make([]float64, 5) // 5 tane float64'den oluşan array

	array := [5]float64{1, 2, 3, 4, 5}
	e := array[0:3]  // [1 2 3]

	fmt.Println(e)


	var toplam float64 = 0

	/*
	for da kullanılan _ boş tanımlayıcı olarak kullanılmaktadır. geri döndürülecek olan değeri tanımlamak için
	kullanılmaktadır. for each if içinde kullanılıyor.
	*/
	for _, eleman := range x {
		toplam += eleman
	}



	fmt.Println("Toplam: ", toplam)
}