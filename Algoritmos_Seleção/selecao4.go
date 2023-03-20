package main

import "fmt"

//Faça um algoritmo que leia a altura e o sexo de uma pessoa
// e mostre se ela está abaixo, dentro ou acima do peso ideal,
//utilizando as fórmulas do exercício 3 da lista de algoritmos sequenciais.

func main() {

	var alt float64
	var sexo string

	fmt.Print("Informe sua aultura: ")
	fmt.Scan(&alt)

	fmt.Print("Seu sexo é masculino ou feminino? ")
	fmt.Scan(&sexo)
	
}
