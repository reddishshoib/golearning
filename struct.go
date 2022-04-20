// Creating Struct in golang and using them accordingly

package main

import (
	"fmt"
)

func main() {
	fmt.Println("structs in golang ")

	shoib := user{"Shoib", "reddish@dev", true, 16}
	fmt.Println(shoib)
	fmt.Printf("Shoib is %v", shoib.Email)
}

type user struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
