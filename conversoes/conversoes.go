package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	notas := 6.9
	notafinal := int(notas)
	fmt.Println(notafinal)

	// cuidado...
	fmt.Println("Teste " + string(97))

	// int para string
	fmt.Println("Teste " + strconv.Itoa(123))

	// string para int
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println("Verdadeiro")
	}

}
