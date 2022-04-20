// This program is to learn use of Defer in golang

package main

import "fmt"

func main() {
	defer fmt.Println("WOrlds")
	fmt.Println("Hello")
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}

}
