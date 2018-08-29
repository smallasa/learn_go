package main

import "fmt"

func main() {
    s1 := make([]int, 5)
    fmt.Printf("The length of s1: %d\n", len(s1))
    fmt.Printf("The capacity of s1: %d\n",cap(s1))
    fmt.Printf("The value os s1: %d\n", s1)

    fmt.Println()

    s2 := make([]int, 5, 8)
    fmt.Printf("The length of s2: %d\n", len(s2))
    fmt.Printf("The capacity of s2: %d\n",cap(s2))
    fmt.Printf("The value os s2: %d\n", s2)
}