package main

import "fmt"

func main() {

	var salario uint64
	fmt.Print("Informe o valor do seu salário: ")
	fmt.Scan(&salario)

	aumento := (salario * 15)
	aumento1 := aumento / 100
	resultado := salario + aumento1

	println("Agora seu novo salário é: ", resultado)

}
