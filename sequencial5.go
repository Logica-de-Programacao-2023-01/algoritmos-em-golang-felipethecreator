package main

import "fmt"

func main() {

	var idade int
	fmt.Print("Informe sua idade: ")
	fmt.Scan(&idade)

	calculo := 365 * idade

	println("Sua idade em dias é: ", calculo)

}
