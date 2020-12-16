# introduction of go

## intro and convenctions 

file `main.go`
```go
package main

import "fmt"

func main() {
	fmt.Println("hello world")
}
```

execute:   
    `go run main.go`

## variables

```go
	var name string = "Username" 
	var version float32 = 1.2
	var age int = 20
```

or with type inference

```go
    var name = "Username"
    name := "Username"
```

## input

```go
    var age int
    fmt.Scanf("%d", &age) // like C
    fmt.Scan(&age) // type-inference
```
