package main

import "fmt"

func main() {

	fmt.Println("numero divisiveis por 3 de 1-100:")
	for numeros := 0; numeros <= 100; numeros++ {
		if numeros%3 == 0 {
			fmt.Println(numeros)
		}
	}
}
