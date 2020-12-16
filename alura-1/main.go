package main

import "fmt"

func main() {
	var age int
	fmt.Scan(&age)

	fmt.Println("age is ", age)
	fmt.Println("variable at mem: ", &age)
}
