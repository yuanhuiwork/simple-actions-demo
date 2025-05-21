package main

import "fmt"

func Sum(a, b int) int {
    return a + b
}

func main() {
    result := Sum(5, 3)
    fmt.Printf("5 + 3 = %d\n", result)
}
