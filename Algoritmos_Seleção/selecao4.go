//Faça um algoritmo que leia a altura e o sexo de uma pessoa e mostre se ela está abaixo,
//dentro ou acima do peso ideal,
//utilizando as fórmulas do exercício 3 da lista de algoritmos sequenciais.

package main

import "fmt"

func main() {

	var sexo string
	var peso, altura float64

	fmt.Print("Informe sua altura: ")
	fmt.Scan(&altura)
	fmt.Print("Informe seu peso: ")
	fmt.Scan(&peso)
	fmt.Print("Informe o seu sexo, sendo 1 para masculino e 2 para feminino: ")
	fmt.Scan(&sexo)
	i := peso / (altura * altura)

	if sexo == "1" && i > 24.9 {
		fmt.Println("Você está acima do peso.")
	} else if sexo == "1" && i < 20 {
		fmt.Println("Você está abaixo do peso.")
	} else if sexo == "1" && i >= 20 && i <= 24.9 {
		fmt.Println("Você está no seu peso ideal.")
	} else if sexo == "2" && i > 23.9 {
		fmt.Println("Você está acima do peso.")
	} else if sexo == "2" && i < 19 {
		fmt.Println("Você está abaixo do peso.")
	} else if sexo == "2" && i >= 19 && i <= 23.9 {
		fmt.Println("Você está no seu peso ideal.")
	} else {
		fmt.Println("Deu não fi...")
	}
}
