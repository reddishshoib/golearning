// Understanding the uses of funcion in golang

package main

import "fmt"

func main() {
	fmt.Println(adder(2, 6, 3))
}
func adder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}
