package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type nextCH chan int

func (ch nextCH) handler(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "<h1> you got %d </h1> ", <-ch)


    // nextID++ unsafe

}

func counter ( ch chan <- int){
        for i:=0 ;;i++{
        time.Sleep(time.Second)
       ch<-i
    }
}

func main (){


    var nextID nextCH = make(chan int)


    go counter(nextID)
    http.HandleFunc("/",nextID.handler)
    log.Fatal(http.ListenAndServe(":3000" , nil))

}
/*
cant write to a chan if somebody  isnt reading it OBV


*
