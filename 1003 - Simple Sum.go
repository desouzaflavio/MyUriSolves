package main

import "fmt"

func main() {

	a := 0
	b := 0

	fmt.Println("Entre com o primeiro valor:")
	fmt.Scanln(&a)

	fmt.Println("Entre com o segundo valor:")
	fmt.Scanln(&b)

	fmt.Println("SOMA = ", sum(a, b))

}

func sum(x int, y int) int {
	return x + y
}
