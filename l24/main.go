package main

import (
	"fmt"
	"time"
)

func main() {
	chans := []chan int{
		make(chan int),
		make(chan int),
	}
	for i := range chans {
		go func(i int, ch chan<- int) {
			for {
				time.Sleep(time.Duration(i) * time.Second)
				ch <- i
			}
		}(i+1, chans[i])
	}

	for i := 0; i < 12; i++ {
		select {
		case val1 := <-chans[0]:
			fmt.Println("val1 ", val1)
		case val2 := <-chans[1]:
			fmt.Println("val1 ", val2)

		}
	}
}

/*
select
 a


*/
