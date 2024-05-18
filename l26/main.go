package main

import "fmt"

func main() {
    ch:= make(chan int ,1)

    ch <- 1

    b, ok := <-ch
    fmt.Println(b,ok)

    close(ch)//the code bellow only works in buffered chanals

    c, ok := <-ch
    fmt.Println(c,ok)



}i

/*
nil chanals get skipped in select


*/
