//Faça um algoritmo que leia vários números inteiros e
//mostre o maior deles. A leitura deve ser interrompida quando for lido o valor zero.

package main

import "fmt"

func main() {
	var numero, maior int

	fmt.Println("Digite vários números inteiros. A leitura será interrompida quando for digitado 0.")

	for {
		fmt.Print("Digite um número inteiro: ")
		fmt.Scan(&numero)

		if numero == 0 {
			break
		}

		if numero > maior {
			maior = numero
		}
	}

	if maior != 0 {
		fmt.Println("O maior número digitado foi:", maior)
	} else {
		fmt.Println("Não foi digitado nenhum número válido.")
	}
}
