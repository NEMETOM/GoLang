package main

import (
	"fmt"
        "strconv"
	"strings"
	"os"
)

type Base struct {
	a int
	b int
	c interface{}
}

func ConcatenateData (b Base) (msg string) {
    
	values := []string{}
	values = append(values, "[1]: " + strconv.Itoa(b.a))
	values = append(values, "[2]: " + strconv.Itoa(b.b))

	//Switch cases here
	switch v := b.c.(type) {
	case int:
		values = append(values, "[3]: " + strconv.Itoa(v))
	case string:
		values = append(values, "[3]: " + v)
	default:
		fmt.Println("This value is not recognized, please review.")
		//Exit the program
		os.Exit(2)
		}
	msg = strings.Join(values, "\n")
	return msg
}


func main() {
	
	//[1] Case Bool - Invalid value - the program will exit
	//e := Base{a:1, b:2, c: true}

	//[2] Case string - Valid value
	//e := Base{a:1, b:2, c: "string value"}

	//[3] Case int - Valid value
	e := Base{a:1, b:2, c: 548}

	//fmt.Println(e)
	fmt.Println(ConcatenateData(e))
}
