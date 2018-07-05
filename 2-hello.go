package main

import "fmt"

func main ()  {
	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[0])
	fmt.Println("Hello" + " World")
	fmt.Println("This is line1\nLine2")
	/*
		Burada farklı olan şey string[index] dediğimizde indexdeki karakterin byte karşılığının sonucu vermesi.
	*/
}