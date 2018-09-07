package main

func main(){
    ch1 := make(chan int, 1)
    ch1 <- 1
    //ch1 <- 2  //通道已满，会造成阻塞

    ch2 := make(chan int, 1)
    //elem, ok := <- ch2  //通道已空，会造成阻塞
    //a, b = elem, ok
    ch2 <- 1

    var ch3 chan int
    //ch3 <- 1  //通道的值为nil，会造成永久阻塞
    //<- ch2    //通道的值为nil，会造成永久阻塞
    _ = ch3
}