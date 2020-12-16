package main

import (
	"fmt"
	"net/http"
)

func main() {
	var numbers [4]int
	numbers[0] = 1
	numbers[1] = 2
	fmt.Println(numbers) // [1 2 0 0]

	slice := []int{1, 2}                       // slice - abstraction of array
	fmt.Println(slice, len(slice), cap(slice)) // [1 2] 2 2
	slice = append(slice, 3)                   // append and duplicate capacity
	fmt.Println(slice, len(slice), cap(slice)) // [1 2 3] 3 4
}

func checkSites() {
	url := "https://www.alura.com.br"
	status := healthCheck(url)
	fmt.Println("the site of url:", url, "is", status)

	url = "https://random-status-code.herokuapp.com"
	status = healthCheck(url)
	fmt.Println("the site of url:", url, "is", status)
}

func healthCheck(site string) string {
	resp, _ := http.Get(site)
	if resp.StatusCode == 200 {
		return "online"
	}
	return "offline"
}
