//Faça um algoritmo que leia dois números inteiros e mostre o resultado da multiplicação entre eles,
//se ambos forem positivos; ou a soma, se pelo menos um deles for negativo.

package main

import "fmt"

func main() {

	var num1, num2 float64
	fmt.Print("Digite um número: ")
	fmt.Scan(&num1)
	fmt.Print("Digite outro número: ")
	fmt.Scan(&num2)

	if num1 > 0 && num2 > 0 {
		fmt.Print("Seu número: ", num1*num2)
	} else if num1 < 0 {
		fmt.Print("Seu número: ", num1+num2)
	} else if num2 < 0 {
		fmt.Print("Seu número: ", num1+num2)
	} else {
		fmt.Print("Seus números não são válidos;")
	}
}
