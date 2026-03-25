package main

import (
	"fmt"
)

const usdEur float64 = 0.87
const usdRub float64 = 83.87

func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

// EUR в RUB
const rubEur float64 = usdRub / usdEur

func check(number float64, currency1 string, currency2 string) bool {
	currency := []string{"USD", "EUR", "RUB"}
	if contains(currency, currency1) && contains(currency, currency2) && number > 0 {
		return true
	} else {
		return false
	}
}

func dataReading() (float64, string, string) {
	var currency1, currency2 string
	var number float64

	for {
		fmt.Printf("Введите название исходной валюты (USD,EUR,RUB)= ")
		fmt.Scan(&currency1)
		fmt.Printf("Введите название конечной валюты (USD,EUR,RUB)= ")
		fmt.Scan(&currency2)
		fmt.Printf("Введите количество для конвертации (Например: 10.0)= ")
		fmt.Scan(&number)
		if check(number, currency1, currency2) {
			return number, currency1, currency2
		} else {
			fmt.Println("Вы ввели неправильные данные, попробуйте ещё раз.")
		}
	}
}

func сurrencyСonversion(number float64, currency1 string, currency2 string) float64 {
	var result float64

	if currency1 == "USD" && currency2 == "RUB" {
		result = number * usdRub
		return result
	} else if currency1 == "USD" && currency2 == "EUR" {
		result = number * usdEur
		return result
	} else if currency1 == "EUR" && currency2 == "RUB" {
		result = number * rubEur
		return result
	} else if currency1 == "EUR" && currency2 == "USD" {
		result = number / usdEur
		return result
	} else if currency1 == "RUB" && currency2 == "USD" {
		result = number / usdRub
		return result
	} else {
		result = number / rubEur
		return result
	}

}

func main() {
	number, currency1, currency2 := dataReading()

	result := сurrencyСonversion(number, currency1, currency2)

	fmt.Printf("%.2F %v = %.2f %v", number, currency1, result, currency2)
}
