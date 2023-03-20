//Faça um algoritmo que leia vários números inteiros e mostre a média aritmética entre eles.
//A leitura deve ser interrompida quando for lido o valor zero.

package main

import "fmt"

func main() {
	var soma, contagem, num int

	fmt.Println("Digite vários números inteiros (digite 0 para sair):")

	for {
		fmt.Scanln(&num)

		if num == 0 {
			break
		}

		soma += num
		contagem++
	}

	if contagem > 0 {
		avg := float64(soma) / float64(contagem)
		fmt.Printf("A média aritmética é: %.2f\n", avg)
	} else {
		fmt.Println("Nenhum número foi digitado.")
	}
}
