package main

import "fmt"

// Function Declearation
// func functionname(parametername type) returntype {
// function body
// }

func calculateBill(price, no int) int {
    totalPrice := price * no
    return totalPrice
}

// Multiple return values
func rectProps(length, width float64) (float64, float64) {
    var area = length * width
    var perimeter = (length + width) * 2
    return area, perimeter
}

// Named return values
func rectProps1(length, width float64) (area, perimeter float64) {
    area = length * width
    perimeter = (length + width) * 2
    return
}

func main() {
    price, no := 96, 6
    totalPrice := calculateBill(price, no)
    fmt.Println("Total price is:", totalPrice)

    area, perimeter := rectProps(10.8, 5.6)
    fmt.Printf("Area %f Perimeter %f\n", area, perimeter)

    area1, perimeter1 := rectProps1(10.8, 5.6)
    fmt.Printf("Area %f Perimeter %f\n", area1, perimeter1)

    // Blank identifier
    area2, _ := rectProps(10.8, 5.6)
    fmt.Println("Area:", area2)
}