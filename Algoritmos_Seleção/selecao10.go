package main

import "fmt"

func main() {

	var idade int64

	fmt.Print("Informe a sua idade: ")
	fmt.Scan(&idade)

	if idade <= 9 && idade > 0 {
		fmt.Println("Mirim.")
	} else if idade >= 10 && idade < 13 {
		fmt.Println("Infantil.")
	} else if idade >= 13 && idade < 17 {
		fmt.Println("Juvenil.")
	} else if idade >= 18 && idade < 130 {
		fmt.Println("Adulto.")
	} else if idade < 0 && idade >= 130 {
		fmt.Println("Mentiroso.")
	}
}
