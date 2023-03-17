package main

import "fmt"

func main() {

	var n1 int64
	fmt.Print("Escreva algum número inteiro: ")
	fmt.Scan(&n1)

	antecessor := n1 - 1
	sucessor := n1 + 1

	fmt.Println("O algarismo antecessor de", n1, "é", antecessor)
	fmt.Println("Seu sucessor é:", sucessor)
}
