package main

import (
	"errors"
	"fmt"
)

const usdEur float64 = 0.87
const usdRub float64 = 83.87

// EUR в RUB
const rubEur float64 = usdRub / usdEur

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

func сurrencyСonversion(number float64, currency1 string, currency2 string) (float64, error) {
	var result float64
	var err error

	if currency1 == "USD" && currency2 == "RUB" {
		result = number * usdRub
	} else {
		err = errors.New("Вы ввели неправильные данные, попробуйте ещё раз.")
	}

	return result, err
}

func main() {
	// Получение данных
	number, currency1, currency2 := dataReading()
	fmt.Printf("%v в %v = %.2f \n", currency1, currency2, number)

	//Example
	result, err := сurrencyСonversion(number, currency1, currency2)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Из %v в %v = %.2f", currency1, currency2, result)
	}
}
