package main

import "fmt"

func main() {

	var n1 float64
	fmt.Print("Digite um número: ")
	fmt.Scan(&n1)

	fmt.Println("O número escolhido foi:", n1, ", seu dobro é:", n1*2, ", seu triplo é:", n1*3,
		"e seu quadruplo é:", n1*4, ".")
}
