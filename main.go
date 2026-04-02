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
	if contains(currencyBase, currency) {
		return true
	} else {
		return false
	}
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
		fmt.Printf("Введите количество для конвертации, введенное число должнобыть положительным. (Например: 10.0)= ")
		fmt.Scan(&number)
		if number > 0 {
			return number
		} else {
			fmt.Printf("Вы ввели %v, это не соответствует предложенным варианту, попробуйте ещё раз.\n", number)
		}
	}
}

type Converter func(float64) float64

var rates = map[[2]string]Converter{
	{"USD", "RUB"}: func(x float64) float64 { return x * usdRub },
	{"USD", "EUR"}: func(x float64) float64 { return x * usdEur },
	{"EUR", "RUB"}: func(x float64) float64 { return x * rubEur },
	{"EUR", "USD"}: func(x float64) float64 { return x / usdEur },
	{"RUB", "USD"}: func(x float64) float64 { return x / usdRub },
	{"RUB", "EUR"}: func(x float64) float64 { return x / rubEur },
}

func Convert(number float64, currency1, currency2 string) float64 {
	if f, ok := rates[[2]string{currency1, currency2}]; ok {
		return f(number)
	}
	return number
}

const usdEur float64 = 0.87
const usdRub float64 = 83.87

// EUR в RUB
const rubEur float64 = usdRub / usdEur

func main() {
	number := createNumber()
	currency1 := createCurrency1()
	currency2 := createCurrency2()

	result := Convert(number, currency1, currency2)

	fmt.Printf("%.2F %v = %.2f %v", number, currency1, result, currency2)
}
