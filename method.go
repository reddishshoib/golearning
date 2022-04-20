// Impementing methods in golang

package main

import (
	"fmt"
)

func main() {
	fmt.Println("structs in golang ")
	shoib := user{"Shoib", "reddish@dev", true, 16}
	shoib.GetStatus()
}

type user struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u user) GetStatus() {
	fmt.Println("Status of user is : ", u.Status)
}
