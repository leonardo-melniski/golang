package main

import "fmt"

func main() {
	number := returnInt()
	fmt.Println(number)

	number = privateFunction(number)

	PublicFunction(number)
}

func returnInt() int {
	return 1
}

func privateFunction(number int) int {
	return number + 1
}

func multipleReturns() (string, int) {
	return "str", 1
}
