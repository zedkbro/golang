package main

import (
	"fmt"
)

func main(){
    var firstName string = "Zinabu";
    var lastName = "Kalayu"
    fmt.Printf("%v\n", firstName)
    fmt.Printf("%T\n ", lastName)
    interns  := []string{"Zinabu", "Tsegay","Daniel"};
    fmt.Println(interns)

    // slice literal
    var mySlice = [5] int {}
    var mySlice2 = interns[0:2]

    fmt.Println(mySlice2)

    x := 3

    switch x {
    case 1:
        fmt.Println("One")
    case 2:
        fmt.Println("Two")
    case 3:
        fmt.Println("Three")
    }

    var  num = 0
    fmt.Println("Enter input")
    fmt.Scanln(&num)

    fmt.Println(num)
    hello()
     }

   func hello(){
     fmt.Println("Hello")
}