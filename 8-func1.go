package  main

import "fmt"

func main(){
	xs := []float64{98,22,54,54,3,4}
	fmt.Println(sum(xs))
	fmt.Println(sumMultiReturn(xs))
	fmt.Println(variadicFunclastNumb(1,2,3,4,5,6,7,8))
}

/*
 Toplam hesabını fonksiyon haline getirelim
*/

func sum(x[]float64) float64{
	total:= 0.0

	for _, v:= range x{
		total += v
	}
	return total
}

/*
Fonksiyonlar birden fazla değer return edebilir.
*/
func sumMultiReturn(x[] float64)(float64,string){
	total := 0.0
	for _, s := range x{
		total	 += s
	}
	return total, "Toplam"
}

/*
Fonksiyon 0 veya daha fazla parametre alabilir bu tip fonksiyonlara variadic fonksiyon deniyor.
*/
func variadicFunclastNumb(numbers ...int)int{
	return numbers[len(numbers) -1]

}

