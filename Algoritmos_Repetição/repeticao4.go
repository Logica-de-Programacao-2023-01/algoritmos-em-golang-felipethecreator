package main // Faça um algoritmo que imprima os múltiplos de 3 de 0 a 30.
import "fmt"

func main() {
	for i := 0; i <= 30; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}
