package main // Faça um algoritmo que imprima a tabuada de
import "fmt"

// multiplicação de 1 a 10 para um número fornecido pelo usuário.

func main() {

	var numero int

	fmt.Print("Digite um número: ")
	fmt.Scan(&numero)

	for i := 1; i <= 10; i++ {
		fmt.Println(i, "X", numero, "=", i*numero)
	}
}
