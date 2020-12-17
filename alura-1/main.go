package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const delay = 2
const filename = "sites.txt"

func main() {
	file, _ := os.Open(filename)
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		fmt.Println(line)

		if err == io.EOF {
			break
		}
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
