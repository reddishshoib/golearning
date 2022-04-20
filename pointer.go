//learn uses of pointer and
// using  & and * in golang

package main

import "fmt"

func main() {
	// fmt.Println("Welcome to pointer world")
	// // var ptr *int
	// mynum := 23
	// var ptr = &mynum
	// *ptr = *ptr + 2
	// fmt.Println("Value of pointer is ", ptr)
	// fmt.Println("Value of pointer is ", mynum)
	ch := make(chan int)
	close(ch)
	val := <-ch
	fmt.Println(val)

}
