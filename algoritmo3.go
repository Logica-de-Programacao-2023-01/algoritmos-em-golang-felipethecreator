package main

import "fmt"

func main() {

	var peso float64
	fmt.Print("Informe seu peso: ")
	fmt.Scan(&peso)

	var altura float64
	fmt.Print("Informe a sua altura: ")
	fmt.Scan(&altura)

	imc := peso / (altura * altura)

	fmt.Print("O valor do seu IMC é: ", imc)

}
