package main

import (
	"fmt"
	"path/filepath"
)

//composition is not inhgerritence

    type Pair struct {
        Path string
        Hash string
    }

    type PairWithLenght struct {

        Pair // this is called promotion
        Length int
    }



func (p PairWithLenght) String() string {
     return   fmt.Sprintln(p.Hash, p.Path,p.Length)

}


func (p Pair) String() string {
     return   fmt.Sprintln(p.Hash, p.Path)

}
//Pair can be passed but PairWithLenght cannot
func (p Pair) Filename () string {
    return filepath.Base(p.Path)
}


type Filenamer interface{
    Filename() string
}

func main (){

    p := Pair{"/usr" , "0x0das"}
    pl := PairWithLenght{Pair{"/usr/bin", "0xdsaf"} , 123}//acception
    fmt.Println(pl.Path,pl.Length) // you can say that the struct Pair has ben embeded into PairWithLenght
    fmt.Print(p)
    //it looks for the method promoted to Pait with length
    fmt.Print(pl)//notise that the string method has been promoted


    fmt.Println(p.Filename())

    var fn Filenamer =  PairWithLenght{Pair{"/usr/bin", "0xdsaf"} , 123}
    fmt.Println(fn)
    fmt.Println(fn.Filename())
}
