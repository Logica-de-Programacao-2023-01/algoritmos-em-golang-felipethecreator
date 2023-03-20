// Faça um algoritmo que leia um número inteiro positivo e mostre todos os seus divisores.

package main

import "fmt"

func main() {

	var numero int
	fmt.Print("Digite um número: ")
	fmt.Scan(&numero)

	fmt.Printf("Os divisores de %d são:", numero)

	for i := 1; i <= numero; i++ {
		if numero%i == 0 {
			fmt.Printf("%d ", i)
		}
	}
}
