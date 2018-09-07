package main

import "fmt"

func main(){
    c1 := make(chan int, 3)
    fmt.Printf("缓冲通道未存入元素时的长度是：%d, 容量是：%d \n", len(c1), cap(c1))
    c1 <- 1
    c1 <- 2
    c1 <- 3
    fmt.Printf("缓冲通道已经存满元素时的长度是：%d, 容量是：%d \n", len(c1), cap(c1))
}