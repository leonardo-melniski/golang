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

multiple returns

```go
func multipleReturns() (string, int) {
	return "str", 1
}
```

example of request

```go
import "net/http"
resp, err := http.Get(site)
```

## array and slice

```go
var array [4]int // array - fixed size
array[0] = 1
array[1] = 2
fmt.Println(array)       // [1 2 0 0]
reflect.TypeOf(array)    // [4]int

slice := []int {1, 2, 3} // slice - abstraction of array
cap(slice)               // 3 - capacity of array
slice = append(slice, 4) // append 4 to slice, and duplicate capacity of inital array
reflect.TypeOf(slice)    // []int
len(slice)               // 4
cap(slice)               // 6 - duplicate initial len
```