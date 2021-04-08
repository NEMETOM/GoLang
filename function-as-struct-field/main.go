package main

import (
	"fmt"
	"strings"
)

//Feeder is used as a data-feed in Author.ConcResult
type Feeder struct {
	FirstName   string
	LastName    string
	Email       string
}

//Author uses function as a struct field
type Author struct {
	ConcResult   func(f Feeder) string
}


func main() {

	feedervalues := Feeder{
		FirstName:     "Tom",
		LastName:      "Nemeth",
		Email:         "tomas.nemeth@dataddo.com",
	}

	result := Author{
		ConcResult:  func(f Feeder) string {
			values := []string{}
			values = append(values, f.FirstName)
			values = append(values, f.LastName)
			values = append(values, f.Email)
            //
			msg := strings.Join(values, "\n")
			//fmt.Println(msg)
			return msg
		},
	}

	fmt.Println(result.ConcResult(feedervalues))
	
}
