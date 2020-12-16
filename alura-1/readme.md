# introduction of go lang

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

## conditions

```go
    if number == 1 {
        fmt.Println("one")
    } else if number == 2 {
        fmt.Println("two")
    } else {
        fmt.Println("fail")
    }
```

```go
	switch number {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("fail")
	}
```
ps: `break` is automatically in go

## function

```go
    func returnInt() int {
        return 1
    }
```
`visibility` by first char 
- `lower` is `p`rivate
- `Upper` is `P`ublic

```go
    func privateFunction() {   
    }
    
    func PublicFunction() {   
    }
```