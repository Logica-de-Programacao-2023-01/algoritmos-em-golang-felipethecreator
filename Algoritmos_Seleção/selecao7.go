//Faça um algoritmo que leia o salário de um funcionário e calcule o seu novo salário com um aumento de 10%
//se o salário for menor que R$ 1000,00; ou de 5% se o salário for maior ou igual a R$ 1000,00.

package main

import "fmt"

func main() {

	var salario float64
	fmt.Print("Informe o valor de seu salário: ")
	fmt.Scan(&salario)

	aumento1 := salario * 1.1
	aumento2 := salario * 1.05

	if salario <= 1000 {
		fmt.Print("Seu novo salário é: ", aumento1)
	} else if salario >= 1000 {
		fmt.Print("Seu nobo salário é: ", aumento2)
	}
}
