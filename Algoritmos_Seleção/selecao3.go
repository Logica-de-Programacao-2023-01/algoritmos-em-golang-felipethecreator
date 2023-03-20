package main // Faça um algoritmo que leia um número inteiro e mostre se ele é par ou ímpar.
import "fmt"

func main() {

	var n1 int
	fmt.Println("Digite um número: ")
	fmt.Scan(&n1)

	if n1%2 == 0 {
		println("Este número é par.")
	} else {
		println("Este número é impar.")
	}
}
