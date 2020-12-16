package main

import (
	"fmt"
	"net/http"
)

func main() {
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
