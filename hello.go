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
    var mySlice = [3] int {1,2,3}
    fmt.Println(mySlice)
    // Q1. how is capacity calculated 
    var mySlice2 = [] int {1, 2, 3, 4, 5, 6}
    var makeSlice = make([]int , 4,10)
     
    // an array with defined length can't be a slice
    fmt.Println(append(mySlice2, 4,5))
    fmt.Println("Capacity of slice ",cap(makeSlice))
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