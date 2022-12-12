package main

import "fmt"

func main() {
	fmt.Println("jogos dos multiplos")
	fmt.Println("multiplos de 3 e de 5")
	for numeros := 1; numeros <= 100; numeros++ {

		if numeros%3 == 0 && numeros%5 == 0 {
			fmt.Println("Pin - Pan")
			continue
		}
		if numeros%5 == 0 {
			fmt.Println("Pan")
			continue
		}

		if numeros%3 == 0 {
			fmt.Println("Pin")
			continue
		}

		fmt.Println(numeros)

	}
}
