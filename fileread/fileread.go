package main

import (
	"fmt"
	"os"
	"io/ioutil"
)

func main() {
	fmt.Println(os.Args)
	//os.Args[1] - this allows user to provide the text file as an argument
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println(string(data))
}

//go run main.go myfile.txt
