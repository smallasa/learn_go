package main

import "fmt"

func main(){
    c1 := make(chan int)
    fmt.Printf("非缓冲通道的长度是：%d, 容量是：%d \n", len(c1), cap(c1))
}