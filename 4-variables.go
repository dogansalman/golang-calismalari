package main

import "fmt"

/*/
	Değişken tanımlamaları
	Mutlaka harf ile başlamalı,
	Harf, sayı, alt çizgi (underscore - _) olabilir,
	camelCase, BumpyCaps, mixedCase diye adlandırılan değişken yazım stilleri de kullanılabilir.
 */

var x string = "Hello World"
var y string


/*
Sabit değerler için  const kullanılır
*/
const s string = "Merhaba Sabit"


/*
Çoklu değişken tanımlama
*/
var i, j int = 1, 2

var (
	a = 5
	b = 10
	c = 15
)


func main(){

	y = "Hello World 2"
	x += y

	// s = "değiştirilemez."

	fmt.Println(x)
	fmt.Println(i, j)
	fmt.Println(s)


	fmt.Println(a,b,c)



}