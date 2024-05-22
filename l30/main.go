package main

import (
	"fmt"
	"sync"
	"time"
)

// deadlock examples
func main() {
	goDeadlock2()
}
func goDeadlock1() {
	ch := make(chan bool)

	go func(ok bool) {
		fmt.Println("Start")
		if ok { //blocking if
			ch <- ok
		}

	}(false)
	<-ch
	fmt.Println("Done")

}
func goDeadlock2() {
	var m sync.Mutex

	done := make(chan bool)

	fmt.Println("Start")
	go func() {
		m.Lock()
		//forget to ulock
	}()

	go func() {
		time.Sleep(1)
		m.Lock()
		defer m.Unlock()

		fmt.Println("Signal")
		done <- true
	}()
	<-done
	fmt.Println("Done")
}
func goDeadlock3() {
	var m1, m2 sync.Mutex

	done := make(chan bool)

	fmt.Println("start")

	go func() {
		m1.Lock()
		defer m1.Unlock()
		time.Sleep(1)
		m2.Lock()
		defer m2.Unlock()
		fmt.Println("signal")
		done <- true
	}()
	go func() {
		m2.Lock()
		defer m2.Unlock()
		time.Sleep(1)
		m1.Lock()
		defer m1.Unlock()
		fmt.Println("signal")
		done <- true
	}()
    //dining philosofers problem
    //always aquire mutexes in order!!!!!!!!!!
	<-done
	fmt.Println("done")
	<-done
	fmt.Println("done")
}
