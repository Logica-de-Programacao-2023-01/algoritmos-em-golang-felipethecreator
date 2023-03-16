package main

import "fmt"

func main() {
	var n1 float64
	fmt.Print("Digite um número: ")
	fmt.Scan(&n1)

	var n2 float64
	fmt.Print("Digite outro número: ")
	fmt.Scan(&n2)

	var n3 float64
	fmt.Print("Digite o ultimo número: ")
	fmt.Scan(&n3)

	fmt.Println("A soma entre esses 3 números é:", n1+n2+n3)

}
