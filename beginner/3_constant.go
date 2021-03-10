package main

import (
	"fmt"
	"math"
)

func main() {
    // Declaring a constant
    const a = 50
    fmt.Println(a)

    // Declaring a group of constants
    const (
        name = "Kai Nguyen"
        age = 18
        country = "Vietnam"
    )
    fmt.Println(name)
    fmt.Println(age)
    fmt.Println(country)
    // name = "Makai" ->  reassignment not allowed

    // The value of a constant should be known at compile time
    var x = math.Sqrt(4) // allowed
    fmt.Println(x)
    // const y = math.Sqrt(5) -> not allowed

    const z bool = false
    fmt.Println(z)
    const gender string = "Male"
    fmt.Println(gender)

}
