package main

import (
    "fmt"
)

var container = []string{"zero", "one", "two"}


func main() {
    container := map[int]string{0:"zero", 1:"one", 2:"two"}
    fmt.Printf("The block is %s.\n", container[1])
}