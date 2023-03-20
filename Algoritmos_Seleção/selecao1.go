package main //Faça um algoritmo que leia dois números inteiros e mostre o maior deles.
import "fmt"

func main() {

	var n1 int64
	var n2 int64

	fmt.Print("Digite um número: ")
	fmt.Scan(&n1)

	fmt.Print("Digite o segundo número: ")
	fmt.Scan(&n2)

	if n1 > n2 {
		fmt.Println("O", n1, "é maior que", n2)
	} else if n2 > n1 {
		fmt.Println("O", n2, "é maior que", n1)
	} else {
		fmt.Println("Ambos os números são iguais!")
	}
}
