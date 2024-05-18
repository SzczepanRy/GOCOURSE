package main

import (
	"fmt"
	"sync"
)

func do() int {

	//could use channals
	//    m := make(chan bool,1 )


    var m sync.Mutex

	var n int64
	var w sync.WaitGroup
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go func() {
            m.Lock()
			//          m <- true

			n++ // data Race
            // or use atomic.AddInt64(&n ,1)
            m.Unlock()
			//          <-m
			w.Done()
		}()
	}
	w.Wait()
	return int(n)
}

func main() {
	//conventional sync

	//mutual exclusion
	fmt.Println(do())

}

//a mutex is better fitted to be a part of a type
//RWMutex
/// !!!! sync.Once guarentees single execytion ex: once.Do(init()) !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

//// sync.Poll provides for efficient and safe reuse of objects ,  but ist a container of interface{}
/* ex:

var bufPool = sync.Pool{
    New:func() interface {}{
        return new(bytes.Buffer)
    }
}

func log(w io.Writer , key ,val string){
    b := bufPool.Get().(*bytes.Buffer)//reflection
    b.Reset()
    //write to it

    w.Write(b.Bytes())
    bufPool.Put(b)

}

*/
