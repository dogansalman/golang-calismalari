package main

import "fmt"

func main(){

		/*
		Go'da Scan, Scanf, ScanIn, ReadString
		Scan: boşluktan sonrasını okumaz.
		ScanIn: boşlukları dahil tüm input değeri algılayabilir yeni bir satıra kadar.
		Scanf: ilk parametresi string değerin formatını belirler. ScanIn ile aynı şekilde çalışır.
		ReadString: Belli bir karaktere ulaşana kadar okumaya devam eder.
		*/



		/*
		Satır satır veri okumanın en iyi yöntemi

		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
		*/



		var input float64
		fmt.Println("Sayı girin:")
		fmt.Scan(&input)

		fmt.Println("Çarpanı girin:")
		var input2 float64
		fmt.Scan(&input2)

		fmt.Println("Sonuç: ", (input * input2))



	/*
	:= bu şekilde tanımlanan değişkenin veri tipini golang algılıyor.
	*/


}