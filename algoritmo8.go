package main

import "fmt"

func main() {

	var dias int
	var valordiaria float64

	fmt.Print("Informe o número de dias de trabalho: ")
	fmt.Scan(&dias)

	fmt.Print("Informe o valor da diária: ")
	fmt.Scan(&valordiaria)

	salario := float64(dias) * valordiaria

	println("O valor do salário do funcionário é:", int64(salario))

}
