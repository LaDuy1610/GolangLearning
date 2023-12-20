package main

import "fmt"

func main() {
	name := "Duy"
	fmt.Println("I Am", name)

	a, b := 1, 2
	fmt.Println("Sum:", a+b, "Mean Value:", (a+b)/2)

	figure := "Circle"
	radius := 5
	pi := 3.14

	fmt.Printf("Radius is %d\n", radius)
	fmt.Printf("Radius is %+d\n", radius)

	fmt.Printf("Pi constant is %f\n", pi)

	fmt.Printf("The diameter of a %s with a Radius of %d is %f\n", figure, radius, float64(radius)*2*pi)

	//%q for quote string
	fmt.Printf("This is a %q\n", figure)

	//%v for replacing any type
	fmt.Printf("The diameter of a %v with a Radius of %v is %v\n", figure, radius, float64(radius)*2*pi)

	//%T to print type
	fmt.Printf("Figure is of %T type\n", figure)

	//%t -> bool
	closed := true
	fmt.Printf("is closed: %t\n", closed)

	//%b ->base 2
	fmt.Printf("%08b\n", 44)

	fmt.Printf("%x\n", 100)

	x := 14.4
	y := 25.6

	fmt.Printf("x * y = %f\n", x*y)
}
