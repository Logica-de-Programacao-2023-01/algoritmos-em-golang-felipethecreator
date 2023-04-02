package main

import "fmt"

func main() {

	var x, y, z float64

	fmt.Print("Digite um número: ")
	fmt.Scan(&x)
	fmt.Print("Digite outro número: ")
	fmt.Scan(&y)
	fmt.Print("Digite mais um número: ")
	fmt.Scan(&z)

	if x > y && y > z {
		fmt.Println(z, ",", y, ",", x)
	} else if x > z && z > y {
		fmt.Println(y, ",", z, ",", x)
	} else if y > x && x > z {
		fmt.Println(z, ",", x, ",", y)
	} else if y > z && z > x {
		fmt.Println(x, ",", z, ",", y)
	} else if z > x && x > y {
		fmt.Println(y, ",", x, ",", z)
	} else if z > y && y > x {
		fmt.Println(x, ",", y, ",", z)
	} else {
		fmt.Println("Não dá")
	}
}
