package main

import "fmt"

func main() {

	var n1 float64
	fmt.Print("Informe um número: ")
	fmt.Scan(&n1)

	var n2 float64
	fmt.Print("Informe o segundo número: ")
	fmt.Scan(&n2)

	var n3 float64
	fmt.Print("Informe o terceiro número: ")
	fmt.Scan(&n3)

	media := (2*n1 + 3*n2 + 5*n3) / (2 + 3 + 5)

	println("Sua média ponderada é:", media)

}
