package main

import (
	"errors"
	"fmt"
)

func main() {
    myFunction("hello", "world", "foo", "bar")
    res,err := divide(12,3)
    fmt.Printf("%v\n",res)
    fmt.Printf("%v",err)

    result := make(chan int)

    // Start a new goroutine to perform a computation
    go func() {
        // Compute the result and send it to the channel
        result <- computeResult()
    }()

    // Receive the result from the channel
    finalResult := <-result
    fmt.Println(finalResult)

    

}
func myFunction(arg1 string, args ...string) {
    fmt.Println("Arg1:", arg1)
    fmt.Println("Variadic args:", args)
}


func divide(a,b int) (int,error){
    if b==0 {
        return 0, errors.New("can't")
    }
    return a/b, nil
}