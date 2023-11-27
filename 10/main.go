package main

import (
	"fmt"
)

func main() {
	total := func() int {
		return sum(1, 2, 3, 45, 567, 67, 88, 9, 5, 6, 34, 52, 3, 4213, 421, 34) * 2
	}()

	fmt.Println(total)
}

func sum(numeros ...int) int {
	total := 0

	for _, n := range numeros {
		total += n
	}

	return total
}
