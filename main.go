package main

import "fmt"

func dataReading() (float64, string, string) {
	var currency1, currency2 string
	var number float64
	fmt.Printf("Введите название исходной валюты (USD,EUR,RUB)= ")
	fmt.Scan(&currency1)
	fmt.Printf("Введите название конечной валюты (USD,EUR,RUB)= ")
	fmt.Scan(&currency2)
	fmt.Printf("Введите количество для конвертации число с точкой (Например: 10.0)= ")
	fmt.Scan(&number)
	return number, currency1, currency2
}

func main() {
	//Получение водимых данных
	number, currency1, currency2 := dataReading()
	fmt.Printf("%v в %v = %.2f \n", currency1, currency2, number)

	const usdEur float64 = 0.87
	const usdRub float64 = 83.87
	//EUR в RUB
	rubEur := usdRub / usdEur
	fmt.Printf("EUR в RUB = %.2f", rubEur)
}
