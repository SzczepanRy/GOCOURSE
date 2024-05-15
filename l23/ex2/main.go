package main

import "fmt"


func generator(limit int ,ch chan<- int ){
    for i:=2; i <limit ; i++{

        ch <- i
    }
    close(ch)
}

func filter(src <- chan int, dst chan<- int , prime int ){
    for i:= range src {
        if i % prime != 0{
            dst <-i
        }
    }
    close(dst)
}

func sive(limit int ){

    ch := make(chan int)

    go generator(limit , ch)

    for  {
            prime , ok := <- ch
            if !ok {break};

        ch1 := make(chan int)

        go filter(ch, ch1, prime)

        ch = ch1
        fmt.Println(prime, " ")
    }
}

func main(){
    sive(100)// 23571117

}
/*
this is a domino effect of channals closing witch allows it to do the prime numbbers ... but i dont realy understand it ......................
but this is not efficient so who cares XD
*/
