package main

import (
	"log"
	"time"
)

func main() {
	const tickRate = 2 * time.Second

    stopper := time.After(5* tickRate)
    ticker := time.NewTicker(tickRate).C

    log.Println("start")
outer:
    for {
        select{
        case <- ticker:
            log.Println("tick")
        case <- stopper:
            break outer
        }
    }


    log.Println("stop")

}
