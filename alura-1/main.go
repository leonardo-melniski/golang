package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const delay = 2
const filename = "sites.txt"
const logFilename = "logs.txt"

func main() {
	openOrCreateFile()
}

func openOrCreateFile() {
	// https://golang.org/pkg/os/#pkg-constants
	file, _ := os.OpenFile(logFilename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)

	// https://golang.org/src/time/format.go
	strTime := time.Now().Format("2006-01-02 15:04:05")
	file.WriteString(strTime + " - example\n")
	file.Close()
}

func openAndReadFile() {
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
