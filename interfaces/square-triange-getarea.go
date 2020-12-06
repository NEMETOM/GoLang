package main
import "fmt"

// new data type Square
type Square struct {
	sideLength float64
}

// new data type Triangle
type Triangle struct {
	base, height float64
}

// A method for type square that returs the calculated area
func (l Square) getArea() float64 {
	return l.sideLength * l.sideLength
}

// A method for type triangle that returs the calculated area
func (t Triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

// Defines methods
type shape interface {
	getArea() float64
}

// This function prints area
func printArea(s shape) {
	fmt.Println(s.getArea())
}


func main() {
	t := Triangle{base: 10, height: 15}
	s := Square{sideLength: 12}

	// variable of type interface
	var a shape

	// Assign to the interface a variable of type "Triangle"
	a = t
	fmt.Println("Area of Triangle is -> ", a.getArea())
	a = s
	// Assign to the interface a variable of type "Square"
	fmt.Println("Area of Square is -> ", a.getArea())
}
