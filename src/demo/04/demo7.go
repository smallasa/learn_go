package main

import (
    "flag"
    "fmt"
)

func main() {
    //方式一
    var name string
    flag.StringVar(&name, "name", "everyone", "The greeting object.")
    flag.Parse()
    //方式二
    var name = *flag.String("name", "everyone", "The greeting object.")
    //方式三
    name := *flag.String("name", "everyone", "The greeting object.")

    fmt.Printf("Hello %s!\n", name)
}
