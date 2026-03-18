package main

import "fmt"

func main() {
	const usdEur float64 = 0.87
	const usdRub float64 = 83.87
	//EUR в RUB
	rubEur := usdRub / usdEur
	fmt.Printf("EUR в RUB = %.2f", rubEur)
}
