package main

import (
	"fmt"
	"net/http"
	"time"
)

const delay = 2

func main() {
	urls := []string{
		"https://www.alura.com.br",
		"https://random-status-code.herokuapp.com",
	}

	for i, url := range urls {
		checkSite(i, url)
		time.Sleep(delay * time.Second)
	}
}

func checkSite(index int, url string) {
	status := healthCheck(url)
	fmt.Println("site[", index, "] :=", url, "is", status)
}

func healthCheck(site string) string {
	resp, _ := http.Get(site)
	if resp.StatusCode == 200 {
		return "online"
	}
	return "offline"
}
