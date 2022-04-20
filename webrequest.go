// Getting response from a given url and saving in a file

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

const url = "https://www.amazon.com"

func main() {
	fmt.Println("This is  Web Request")
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Respones is of %T\n", response)
	defer response.Body.Close()

	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(databytes)
	witeinfile(content)

}

func witeinfile(content string) {
	file, err := os.Create("./index.html")
	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("Length is ", length)
	defer file.Close()
}
