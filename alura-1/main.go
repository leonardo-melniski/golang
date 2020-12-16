package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)

	switch number {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("fail")
	}
}
