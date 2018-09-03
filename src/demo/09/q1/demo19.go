package main

import "fmt"

func main() {
    var badMap2 = map[interface{}] int {
        "1" : 1,
        []int{2} : 2,
        3 : 3,
    }

    fmt.Println(badMap2)
}