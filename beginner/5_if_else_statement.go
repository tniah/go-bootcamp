package main

import "fmt"

func main() {
    num := 10
    if num%2 == 0 {
        fmt.Println("The number", num, "is even")
    }

    num = 11
    if num % 2 == 0 {
        fmt.Println("The number", num, "is even")
    } else {
        fmt.Println("The number", num, "is odd")
    }

    num = 99
    if num <= 50 {
        fmt.Println(num, "is less than or equal to 50")
    } else if num >= 51 && num <= 100 {
        fmt.Println(num, "is between 51 and 100")
    } else {
        fmt.Println(num, "is greater than 100")
    }

    if num1 := 101; num % 2 == 0 {
        fmt.Println(num1, "is even")
    } else {
        fmt.Println(num1, "is odd")
    }
}