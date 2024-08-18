package main

import "fmt"

func main() {
    

    myFunction("hello", "world", "foo", "bar")
}
func myFunction(arg1 string, args ...string) {
    fmt.Println("Arg1:", arg1)
    fmt.Println("Variadic args:", args)
}