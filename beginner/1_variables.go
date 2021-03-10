package main

import "fmt"

func main() {

    // Declaring a single variable
    var age int
    fmt.Println("My age is", age)
    // My age is 0
 
    // Declaring a variable with an initial value
    var age1 int = 18
    fmt.Println("My age is", age1)
    // My age is 18

    // Multiple variable declaration
    var width, height int = 100, 300
    fmt.Println("width is", width, "height is", height)
    // width is 100 height is 300
    var width1, height1 = 100, 300
    fmt.Println("width is", width1, "height is", height1)

    var (
        name = "Kai Nguyen"
        age2 = 18
    )
    fmt.Println("My name is", name, "I am", age2, "years-old")
    // My name is Kai Nguyen. I am 18 years-old

    // Short hand declaration
    count := 10
    fmt.Println("Count", count)
}
