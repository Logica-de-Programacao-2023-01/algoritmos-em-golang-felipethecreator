package main //Faça um algoritmo que leia três números inteiros e mostre o menor deles.
import "fmt"

func main() {

	var n1 int
	var n2 int
	var n3 int

	fmt.Print("Digite um número: ")
	fmt.Scan(&n1)
	fmt.Print("Digite o segundo número: ")
	fmt.Scan(&n2)
	fmt.Print("Digite o terceiro número: ")
	fmt.Scan(&n3)

	if n1 < n2 && n1 < n3 {
		fmt.Println("O", n1, "é o menor número.")
	} else if n2 < n1 && n2 < n3 {
		fmt.Println("O", n2, "é o menor número.")
	} else if n3 < n2 && n3 < n1 {
		fmt.Println("O", n3, "é o menor número.")
	} else {
		fmt.Println("Todos os números são iguais!")
	}
}
