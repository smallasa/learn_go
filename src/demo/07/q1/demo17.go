package main

import "fmt"

func main() {
    s1 := []int{1,2,3,4,5}
    printSlice("s1", s1)

    fmt.Println()

    s1 = shrinkSlice(s1)
    printSlice("s1", s1)
}

func printSlice(s string, x []int){
    fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func shrinkSlice(x []int) []int {
    if (cap(x) > 0) {
        x = x[0:cap(x)-1]
    }
    return x
}