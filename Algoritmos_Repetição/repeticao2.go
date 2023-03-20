package main // Faça um algoritmo que imprima os números pares de 0 a 20.
import "fmt"

func main() {
	for i := 0; i <= 20; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
