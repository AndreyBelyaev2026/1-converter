package main

import (
	"fmt"
)

func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func checkContains(currency string) bool {
	currencyBase := []string{"USD", "EUR", "RUB"}
	return contains(currencyBase, currency)
}

func createCurrency1() string {
	var currency1 string
	for {
		fmt.Printf("Введите название исходной валюты (USD,EUR,RUB)= ")
		fmt.Scan(&currency1)
		if checkContains(currency1) {
			return currency1
		} else {
			fmt.Printf("Вы ввели %v, это не соответствует предложенным вариантам исходной валюты, попробуйте ещё раз.\n", currency1)
		}
	}
}

func createCurrency2() string {
	var currency2 string
	for {
		fmt.Printf("Введите название целевой валюты (USD,EUR,RUB)= ")
		fmt.Scan(&currency2)
		if checkContains(currency2) {
			return currency2
		} else {
			fmt.Printf("Вы ввели %v, это не соответствует предложенным вариантам целевой валюты, попробуйте ещё раз.\n", currency2)
		}
	}
}

func createNumber() float64 {
	var number float64
	for {
		fmt.Printf("Введите количество для конвертации, введенное число должно быть положительным. (Например: 10.0)= ")
		fmt.Scan(&number)
		if number > 0 {
			return number
		} else {
			fmt.Printf("Вы ввели %v, это не соответствует предложенным варианту, попробуйте ещё раз.\n", number)
		}
	}
}

type Converter func(float64) float64

const usdEur float64 = 0.87
const usdRub float64 = 83.87

// EUR в RUB
const rubEur float64 = usdRub / usdEur

func Convert(number float64, currency1, currency2 string, rates *map[[2]string]Converter) float64 {
	if f, ok := (*rates)[[2]string{currency1, currency2}]; ok {
		return f(number)
	}
	return number
}

func main() {
	// Передача указателя на карту
	var rates = map[[2]string]Converter{
		{"USD", "RUB"}: func(x float64) float64 { return x * usdRub },
		{"USD", "EUR"}: func(x float64) float64 { return x * usdEur },
		{"EUR", "RUB"}: func(x float64) float64 { return x * rubEur },
		{"EUR", "USD"}: func(x float64) float64 { return x / usdEur },
		{"RUB", "USD"}: func(x float64) float64 { return x / usdRub },
		{"RUB", "EUR"}: func(x float64) float64 { return x / rubEur },
	}

	number := createNumber()
	currency1 := createCurrency1()
	currency2 := createCurrency2()

	result := Convert(number, currency1, currency2, &rates)

	fmt.Printf("%.2f %s = %.2f %s\n", number, currency1, result, currency2)
}
