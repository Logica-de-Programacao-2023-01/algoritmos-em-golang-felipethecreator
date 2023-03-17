package main

import "fmt"

func main() {

	var peso float64

	fmt.Print("Informe seu peso: ")
	fmt.Scan(&peso)

	conversor := peso * 2.2046

	fmt.Println("Seu peso convertido em libras é:", float64(conversor))

}
