//Faça um algoritmo que leia um número inteiro e mostre se ele é múltiplo de 3 e de 5 ao mesmo tempo.

package main

import "fmt"

func main() {

	var numero1 int
	fmt.Print("Informe um número: ")
	fmt.Scan(&numero1)

	if numero1%3 == 0 && numero1%5 == 0 {
		fmt.Println("Seu número É divisível por 3 e 5.")
	} else {
		fmt.Println("Seu número NÃO É divisível por 3 e 5.")
	}
}
