package main

import "fmt"
func main(){
	// slice of ints from 0 to 10
	var numbs = []int{0,1,2,3,4,5,6,7,8,9,10}

	// loop through the numbers and check if a number in the range is odd or even
	for _, number := range numbs {
		if number%2 == 0 {
			fmt.Printf("%v is Even number\n", number)
		} else {
			fmt.Printf("%v is Odd number\n", number)
		}
	}
}
