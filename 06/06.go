/*
6
Sum square difference
https://projecteuler.net/problem=6
Autorzy: C99 ANSIC
C99 ANSIC - Całość implementacji rozwiązania.
*/

package main
import "fmt"

func main() {
	var sumaKwadratów, sumaCiągu, kwadratSumyCiągu uint64
	var limit uint64 = 100

	//Obliczenie sumy kwadratów
	for i := uint64(1); i <= limit; i++ {
		sumaKwadratów += i*i
	}
	//Obliczenie sumy ciągu
	for i := uint64(1); i <= limit; i++ {
		sumaCiągu += i
	}
	//Obliczenie kwadratuSumyCiągu
	kwadratSumyCiągu = sumaCiągu * sumaCiągu
	//Obliczenie i wyświetlenie wyniku
	fmt.Println(kwadratSumyCiągu-sumaKwadratów)
}