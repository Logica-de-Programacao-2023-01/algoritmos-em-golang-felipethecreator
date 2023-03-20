package main

import "fmt"

func main() {

	var produto float64

	fmt.Print("Informe o valor do produto: ")
	fmt.Scan(&produto)

	desconto := (produto * 10)
	desconto2 := (desconto / 100)
	resultado := (produto - desconto2)

	fmt.Println("O valor do produto com 10% de desconto é:", float64(resultado))
}
