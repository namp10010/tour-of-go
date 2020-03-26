package main

import "fmt"

func main() {
    ch := make(chan int, 2)
    ch <- 100
    ch <- 200
    fmt.Println(<-ch)
    fmt.Println(<-ch)
}
