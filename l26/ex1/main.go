package main

import (
	"fmt"
	"time"
)

type T struct {
	i byte
	b bool
}


func send(i int ,  ch chan<- *T ){
    t := &T{i:byte(i)}
    ch <-t
    t.b = true //unsafe
}


func main() {
    vs := make([]T,5)
    ch := make(chan *T,5) //  ttry unbufered
    for i:= range vs{
        go send(i,ch)
    }

    time.Sleep(time.Second)
    //coppy quick
    for i := range vs{
        vs[i] = *<-ch
    }
    // print later

    for _, v := range vs{
        fmt.Println(v)
    }


}
/*
why buffer
avoid leeks (form abandend chanels)

OCUNTING SEMAGHORE
mits work in progress

*/
