package main

import (
    "fmt"
    "io"
    "os"
)


type ByteCounter int

func (b *ByteCounter) Write(p []byte) (int,error){

    *b += ByteCounter(len(p))

    return len(p),nil

}


func main (){


    var c ByteCounter

    f1 ,_ := os.Open("a.txt")
    f2:= &c

    n, _ := io.Copy(f2, f1)

    fmt.Println("coppied", n , "bytes")

    fmt.Println(c)
}

/*

compatibility

the recever must be of right type (pointer or value )
ex1 :

type IntSet sruct { ... }
func (*IntSet) String() string
var _ = IntSet{}.String will not work because of the lack of address of IntSet

var s IntSet
var _ = s.String works





*/
