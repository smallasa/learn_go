package main

import "fmt"

func main(){
    a1 := [7]int{1, 2, 3, 4, 5, 6, 7}
    fmt.Printf("a1: %v (len: %d, cap: %d)\n", a1, len(a1), cap(a1))

    a9 := a1[1:4]
    fmt.Printf("a9: %v (len: %d, cap: %d)\n", a9, len(a9), cap(a9))

    for i := 1; i <= 5; i++ {
        a9 = append(a9, i)
        fmt.Printf("a9(%d): %v (len: %d, cap: %d)\n", i, a9, len(a9), cap(a9))
    }

    fmt.Printf("a1: %v (len: %d, cap: %d)\n", a1, len(a1), cap(a1))

    fmt.Println()
}