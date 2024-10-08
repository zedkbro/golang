package main

import (
	"fmt"
)


type Animal struct{
    name string
    Age int
}

func main(){
    var firstName string = "Zinabu";
    var lastName = "Kalayu"
    fmt.Printf("%v\n", firstName)
    fmt.Printf("%T type \n", lastName)
    interns:= []string{"Zinabu", "Tsegay","Daniel"};
    fmt.Println(interns)
    //  mySlice5 := append(interns, )
    // slice literal
    var mySlice = [3] int {1,2,3}
    fmt.Println(mySlice)
    // Q1. how is capacity calculated 
    var mySlice2 = [] int {1, 2, 3, 4, 5, 6}
    var makeSlice = make([]int , 4,10)
     
    // an array with defined length can't be a slice
    // var mySlice3 = [6] int {1, 2, 3, 4, 5, 6}

    fmt.Println(append(mySlice2, makeSlice...))
    fmt.Println("Capacity of slice ",cap(makeSlice))
    fmt.Println(mySlice2)
    x:= 3
    var str string
    if x > 3 {
         str = "Hi"
    }else{
        str = "Hello"
    }

    fmt.Println(str)
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
    loop()

    // Go map
    var mapA = map[string] int {"Age":21}
    // var map3 map[string] int; declaring a map - default value is nil
    var mapB = make(map[string] int)

    mapB["Salary"] = 34000;
    fmt.Println(mapA["Age"])
    fmt.Println(mapB)
    delete(mapB, "Salary")

  

    var person  Animal
    var dog Animal

    person.name = "Zinabu"
    dog.name ="Tafro"
    fmt.Println(person.name)
    fmt.Println(dog.name)

    animalFunc(person)

    

}

   func hello(){
     fmt.Println("Hello")
}

func loop(){
    for i := 0;i<5; i++{
    fmt.Print(i)
      }
    var fruits = [3] string {"Banana", "Apple", "Orange"}
    for _,value := range fruits {
        fmt.Println(value)
    }
}

func animalFunc(infant Animal){
    fmt.Println(infant.name)
}
