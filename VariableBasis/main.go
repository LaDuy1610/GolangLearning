package main

import "fmt"

func main() {
	var age int = 30
	fmt.Println("Age: ", age)

	var name string = "Duy"
	//fmt.Println("Your name: ", name)
	_ = name

	s := "Learning Golang!"
	fmt.Println(s)

	car, cost := "BMW", 10000
	fmt.Println(car, cost, name)

	var (
		salary float64
		job    string
		gender bool
	)
	fmt.Println(salary, job, gender)
}
