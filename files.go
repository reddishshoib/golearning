// This program is to read and write in files
// in golang using ioutil
package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// fmt.Println("Welcome to files in golang ")
	// contnt := "This is to be going in files"
	// file, err := os.Create("./createdfile.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// length, err := io.WriteString(file, contnt)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Length is ", length)
	// defer file.Close()
	readFile("./createdfile.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("Text Data inside file is \n", string(databyte))
}
